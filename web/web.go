package web

import (
	"encoding/json"
	"html/template"
	"log"
	"math/big"
	"net/http"
	"path"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/mathieugilbert/aabot/adapters"
	"github.com/mathieugilbert/aabot/cmd/config"
	"github.com/mathieugilbert/aabot/orderbook"
)

// ServiceStore wraps needed services
type ServiceStore struct {
	Config   *config.Configuration
	Ethereum *adapters.Ethereum
	Redis    *redis.Client
	DB       orderbook.Datastore
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Start the web server
func Start(ss *ServiceStore) {
	router := httprouter.New()

	router.GET("/", ss.root)
	router.GET("/monitor", ss.monitor)
	//router.GET("/market/:id", ss.market)
	router.GET("/markets", ss.markets)
	router.GET("/wallets", ss.wallets)
	router.GET("/orderbook", ss.orderbook)
	router.GET("/ws", ss.websocket)

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

	exs, err := ss.DB.ExchangesJoined()
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

/*
func (ss *ServiceStore) market(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	mid, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		log.Printf("Invalid market ID: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type Response struct {
		Market *orderbook.Market `json:"market"`
	}

	m, err := ss.DB.Market(mid)
	if err != nil {
		log.Printf("Error getting market: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := &Response{Market: m}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
*/
func (ss *ServiceStore) wallets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type Response struct {
		Wallets []*orderbook.Wallet `json:"wallets"`
	}

	ws, err := ss.DB.Wallets()
	if err != nil {
		log.Printf("Error getting wallets: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := &Response{Wallets: ws}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (ss *ServiceStore) markets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type Response struct {
		Markets []*orderbook.Market `json:"markets"`
	}

	ms, err := ss.DB.MarketsJoined()
	if err != nil {
		log.Printf("Error getting markets: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := &Response{Markets: ms}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (ss *ServiceStore) orderbook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := r.URL.Query()

	mid, err := strconv.Atoi(q.Get("mid"))
	if err != nil {
		http.Error(w, "Invalid market id", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(q.Get("n"))
	if err != nil {
		http.Error(w, "Invalid count", http.StatusBadRequest)
		return
	}

	s, err := ss.DB.LowestSells(mid, n)
	if err != nil {
		http.Error(w, "Error getting lowest sells", http.StatusInternalServerError)
		return
	}
	b, err := ss.DB.HighestBuys(mid, n)
	if err != nil {
		http.Error(w, "Error getting highest buys", http.StatusInternalServerError)
		return
	}

	var os = make([]*orderbook.Order, 0)

	for _, o := range s {
		os = append(os, o)
	}
	for _, o := range b {
		os = append(os, o)
	}

	type Response struct {
		Orders []*orderbook.Order `json:"orders"`
	}
	resp := &Response{Orders: os}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (ss *ServiceStore) websocket(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		log.Println("rejected 403", r.Header.Get("Origin"))
		http.Error(w, "Origin not allowed", http.StatusForbidden)
		return
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading ws: %v\n", err)
		http.Error(w, "Could not upgrade websocket connection", http.StatusBadRequest)
		return
	}
	defer c.Close()

	type Message struct {
		Message string `json:"message"`
		Value   string `json:"value"`
	}

	for {
		mt, m, err := c.ReadMessage()
		if err != nil {
			log.Printf("WS read error: %v\n", err)
			break
		}

		msg := &Message{}
		if err := json.Unmarshal(m, msg); err != nil {
			log.Printf("Bad message format: %v\n", err)
			break
		}

		switch msg.Message {
		case "wallet":
			err = c.WriteMessage(mt, []byte(msg.Value))
		default:
			err = c.WriteMessage(mt, []byte("Unknown command"))
		}

		if err != nil {
			log.Printf("WS write error: %v\n", err)
			break
		}
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

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
