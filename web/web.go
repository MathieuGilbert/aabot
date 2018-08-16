package web

import (
	"encoding/json"
	"html/template"
	"log"
	"math/big"
	"net/http"
	"path"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/mathieugilbert/aabot/adapters"
	"github.com/mathieugilbert/aabot/cmd/config"
	"github.com/mathieugilbert/aabot/orderbook"
	goed "github.com/mathieugilbert/go-etherdelta"
)

// ServiceStore wraps needed services
type ServiceStore struct {
	Config   *config.Configuration
	Ethereum *adapters.Ethereum
	GoEd     *goed.Service
	Redis    *redis.Client
	DB       orderbook.Datastore
}

// Start the web server
func Start(ss *ServiceStore) {
	router := httprouter.New()

	router.GET("/", ss.root)
	router.GET("/monitor", ss.monitor)
	router.GET("/market/:exchange/:token", ss.market)

	// serve static assets
	router.ServeFiles("/web/js/*filepath", http.Dir("web/js"))
	router.ServeFiles("/web/css/*filepath", http.Dir("web/css"))
	router.ServeFiles("/web/vendor/*filepath", http.Dir("web/vendor"))

	// start the web server
	log.Println("Web server ready.")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func (ss *ServiceStore) root(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type Section struct {
		Token *orderbook.Token
		Buys  []*orderbook.Order
		Sells []*orderbook.Order
	}
	type Data struct {
		DexBotNum *big.Int
		Sections  []*Section
	}
	data := &Data{}

	//num, err := ss.Ethereum.DexBot.Number(nil)
	//if err != nil {
	//	http.Error(w, "Error reading number from DexBot", http.StatusInternalServerError)
	//	return
	//}
	data.DexBotNum = big.NewInt(555)

	ts, err := ss.DB.Tokens()
	if err != nil {
		log.Println("Couldn't get tokens")
		http.Error(w, "Couldn't get tokens", http.StatusInternalServerError)
		return
	}
	for _, t := range ts {
		buys, err := ss.DB.HighestBuys(t.ID, 5)
		if err != nil {
			log.Println("Couldn't find highest buys")
			http.Error(w, "Couldn't find highest buys", http.StatusInternalServerError)
			return
		}

		sells, err := ss.DB.LowestSells(t.ID, 5)
		if err != nil {
			log.Println("Couldn't find lowest sells")
			http.Error(w, "Couldn't find lowest sells", http.StatusInternalServerError)
			return
		}

		data.Sections = append(data.Sections, &Section{
			Token: t,
			Buys:  buys,
			Sells: sells,
		})
	}

	t := pageTemplate("web/templates/root.html.tmpl")
	t.Execute(w, data)
}

func (ss *ServiceStore) monitor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type Data struct {
		Exchanges []*orderbook.Exchange
		Tokens    []*orderbook.Token
	}
	data := &Data{}

	exs, err := ss.DB.Exchanges()
	if err != nil {
		log.Println("Error getting exchanges")
		http.Error(w, "Error getting exchanges", http.StatusInternalServerError)
		return
	}
	data.Exchanges = exs

	tks, err := ss.DB.Tokens()
	if err != nil {
		log.Println("Error getting tokens")
		http.Error(w, "Error getting tokens", http.StatusInternalServerError)
		return
	}
	data.Tokens = tks

	t := pageTemplate("web/templates/monitor.html.tmpl")
	t.Execute(w, data)
}

func (ss *ServiceStore) market(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	e, err := strconv.Atoi(p.ByName("exchange"))
	if err != nil {
		log.Printf("Invalid exchange ID: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t, err := strconv.Atoi(p.ByName("token"))
	if err != nil {
		log.Printf("Invalid token ID: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type Response struct {
		Market *orderbook.Market `json:"market"`
	}

	m, err := ss.DB.MarketByExTok(e, t)
	if err != nil {
		log.Printf("Error getting market: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := &Response{Market: m}

	tkn, err := ss.DB.Token(t)
	if err != nil {
		log.Printf("Error getting token: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := "0xE80864b6B92c1F7186Fb97bbAE0b2098C4CAB014"
	bal, err := ss.Ethereum.EtherDeltaInstance().BalanceOf(&bind.CallOpts{Pending: true}, common.HexToAddress(tkn.Address), common.HexToAddress(user))
	log.Printf("Token: %v\nBalance: %v\n", tkn.Address, bal)
	/*
		var opts = &goed.GetTokenBalanceOpts{TokenAddress: tkn.Address, UserAddress: "0x3e25f0ba291f202188ae9bda3004a7b3a803599a"}
		bal, err := ss.GoEd.GetTokenBalance(opts)
	*/
	if err != nil {
		log.Printf("Error getting balance: %v\ntoken: %v, user: %v\n", err, tkn.Address, user)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Market.Balance = bal.String()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func pageTemplate(ts ...string) *template.Template {
	f := []string{"web/templates/index.html.tmpl"}
	for _, t := range ts {
		f = append(f, t)
	}
	return template.Must(template.New(path.Base(f[0])).Funcs(funcMaps()).ParseFiles(f...))
}

func funcMaps() map[string]interface{} {
	return template.FuncMap{
		"toEth": toEth,
	}
}

func toEth(field string) *big.Float {
	n := new(big.Float)
	f, ok := n.SetString(field)
	if !ok {
		return n
	}

	return new(big.Float).Quo(f, big.NewFloat(1000000000000000000))
}
