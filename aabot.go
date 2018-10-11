package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	_ "github.com/jinzhu/gorm/dialects/postgres" // db driver
	"github.com/mathieugilbert/aabot/adapters"
	"github.com/mathieugilbert/aabot/cmd/config"
	idex "github.com/mathieugilbert/go-idex"

	"github.com/mathieugilbert/aabot/orderbook"
	"github.com/mathieugilbert/aabot/web"
	goed "github.com/mathieugilbert/go-etherdelta"
)

var (
	// Config holds the app configuration
	Config *config.Configuration
	// ConfigFile contains app settings
	ConfigFile = "private/mainnet/config.json"
	// SeedFile contains seed data
	SeedFile = "private/mainnet/seed.json"
)

func init() {
	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	// load config values
	loadConfig(ConfigFile)
	// migrate database schema
	orderbook.Migrate(Config.DBString())
}

func main() {
	// Database init
	db, err := orderbook.NewDB(Config.DBString())
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)

	// seed config data
	db.ClearTempTables()
	db.Seed(SeedFile)

	// Ethereum network
	eth, err := adapters.NewEthereum(Config.Ethereum.ProviderURI)
	if err != nil {
		log.Panicf("Couldn't connect to Ethereum network: %v\n%v\n", Config.Ethereum.ProviderURI, err)
	}

	// Redis client
	r := redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Path,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Wrap needed services
	ss := &web.ServiceStore{
		Config:   Config,
		Ethereum: eth,
		Redis:    r,
		DB:       db,
	}
	/*
		if err := db.SyncBalances(eth, Config.Ethereum.UserAddress); err != nil {
			log.Panicf("Unable to sync wallet balances: %v\n", err)
		}
		if err := db.SyncMarketBalances(eth, Config.Ethereum.UserAddress); err != nil {
			log.Panicf("Unable to sync exchange balances: %v\n", err)
		}
	*/
	es, err := db.Exchanges()
	if err != nil {
		log.Panicf("Unable to retrieve exchanges: %v\n", err)
	}

	// wait for all exchanges to load
	var wg sync.WaitGroup
	wg.Add(len(es))

	for _, e := range es {
		go func(e *orderbook.Exchange, wg *sync.WaitGroup) {
			e.LoadOrderbooks(db)
			wg.Done()
		}(e, &wg)
	}
	wg.Wait()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	defer func() {
		signal.Stop(interrupt)
		cancel()
	}()

	go ConsumeEtherDelta(ctx, db)
	go ConsumeIdex(ctx, db)

	go web.Start(ss)

	select {
	case <-interrupt:
		cancel()
	}
}

func ConsumeEtherDelta(ctx context.Context, db *orderbook.DB) {
	url := "wss://api.forkdelta.com/socket.io/?EIO=3&transport=websocket"
	c, resp, err := websocket.DefaultDialer.Dial(url, nil)
	if err == websocket.ErrBadHandshake {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		log.Printf("%+v\n", buf.String())
		log.Printf("handshake failed with status %d", resp.StatusCode)
	}
	if err != nil {
		log.Panicf("Dial(%v) error: %v\n", url, err)
	}
	defer c.Close()

	e, err := db.ExchangeByName("EtherDelta")
	if err != nil {
		log.Panicf("Can't find EtherDelta in db")
	}

	var respRxp = regexp.MustCompile(`(\d+)(.*)`)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			if err.Error() == "websocket: close 1000 (normal)" {
				log.Println("Websocket closed")
				return
			}
			log.Println("ReadMessage error:", err)
			continue
		}

		match := respRxp.FindStringSubmatch(string(message))
		if len(match) != 3 {
			log.Printf("Regex didn't find matches in: %v", message)
			continue
		}

		switch match[1] {
		case "0":
			log.Printf("Open: %v\n", match[2])
		case "40":
			log.Println("Connected to websocket")
		case "42":
			var orderRxp = regexp.MustCompile(`\[.(.*?).,(.*)]`)
			matchOrder := orderRxp.FindStringSubmatch(match[2])
			if len(matchOrder) != 3 {
				log.Printf("42, not orders: %v\n", match[2])
				break
			}

			switch matchOrder[1] {
			case "orders":
				resp := &goed.OrderBook{}

				if err = json.Unmarshal([]byte(matchOrder[2]), resp); err != nil {
					log.Printf("Can't unmarshal orders: %v\n%v\n", matchOrder[2], err)
					break
				}

				count, errr := e.ProcessNewOrdersEtherDelta(db, resp)
				if errr != nil {
					log.Printf("Unable to process new orders: %v\n", errr)
					break
				}
				if count > 0 {
					log.Printf("Inserted %v new orders\n", count)
				}
			case "trades":
				type Trade struct {
					TxHash     string `json:"txHash"`
					Date       string `json:"date"`
					Price      string `json:"price"`
					Side       string `json:"side"`
					Amount     string `json:"amount"`
					AmountBase string `json:"amountBase"`
					Buyer      string `json:"buyer"`
					Seller     string `json:"seller"`
					TokenAddr  string `json:"tokenAddr"`
				}
				var trades []*Trade

				if err = json.Unmarshal([]byte(matchOrder[2]), &trades); err != nil {
					log.Panicf("Can't unmarshal trades: %v\n%v\n", matchOrder[2], err)
					break
				}

				for _, trade := range trades {
					// re-fetch orderbook for token
					m, err := db.MarketByExTokAddr(e.ID, trade.TokenAddr)
					if err != nil {
						log.Printf("Error fetching market for exchange %v, token %v\n", e.ID, trade.TokenAddr)
						continue
					}

					var wg sync.WaitGroup
					wg.Add(1)
					go func(m *orderbook.Market) {
						m.LoadOrderbookEtherDelta(db, &wg)
					}(m)
					wg.Wait()

					log.Printf("New trades, refreshed orderbook for token %v on exchange %v\n", trade.TokenAddr, e.Name)
				}

			}
		default:
			log.Printf("other message: %v\n", match[0])
		}

		select {
		case <-ctx.Done():
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error writing close message:", err)
				return
			}

			<-time.After(time.Second)
			return
		}
	}
}

func ConsumeIdex(ctx context.Context, db *orderbook.DB) {
	i := idex.New()
	if err := i.Socket.Connect(); err != nil {
		log.Panic(err)
	}
	defer i.Socket.Conn.Close()

	response := make(chan idex.SocketResponse)
	go i.Socket.Monitor(response)

	for {
		select {
		case r := <-response:
			if r.PushCancel != nil {
				if err := db.DeleteOrderByHash(r.PushCancel.Hash); err != nil {
					log.Panicf("Error deleting order by hash %v: %v\n", r.PushCancel.Hash, err)
				}
			}
			if r.TradeInserted != nil {
				fmt.Printf("got a trade: %+v\n", r.TradeInserted)
				// doesn't look like it's possible to update the order, so get the orderbook
				e, err := db.ExchangeByName("IDEX")
				if err != nil {
					log.Panicf("Error fetching IDEX exchange: %v\n", err)
				}

				eth, err := db.TokenByName("ETH")
				if err != nil {
					log.Panicf("Error fetching ETH token: %v\n", err)
				}

				var tknAddr string
				if r.TradeInserted.TokenSell == eth.Address {
					tknAddr = r.TradeInserted.TokenBuy
				} else {
					tknAddr = r.TradeInserted.TokenSell
				}

				m, err := db.MarketByExTokAddr(e.ID, tknAddr)
				if err != nil {
					// token not one we care about
					continue
				}

				var wg sync.WaitGroup
				wg.Add(1)
				go m.LoadOrderbookIDEX(db, &wg)
				wg.Wait()
			}
			if r.OrderInserted != nil {
				var isBuy bool
				var tknAddr string
				var vol string
				var price string

				eth, err := db.TokenByName("ETH")
				if err != nil {
					log.Panicf("Error fetching ETH token: %v\n", err)
				}

				b := new(big.Int)
				if _, ok := b.SetString(r.OrderInserted.AmountBuy, 10); !ok {
					log.Panicf("Error converting AmountBuy to big.Int: %v, %v\n", r.OrderInserted.AmountBuy, err)
				}

				s := new(big.Int)
				if _, ok := s.SetString(r.OrderInserted.AmountSell, 10); !ok {
					log.Panicf("Error converting AmountSell to big.Int: %v, %v\n", r.OrderInserted.AmountSell, err)
				}

				p := new(big.Int)

				if r.OrderInserted.TokenSell == eth.Address {
					isBuy = true
					tknAddr = r.OrderInserted.TokenBuy
					vol = r.OrderInserted.AmountBuy
					p.Div(b, s)
				} else {
					isBuy = false
					tknAddr = r.OrderInserted.TokenSell
					vol = r.OrderInserted.AmountSell
					p.Div(s, b)
				}

				price = fmt.Sprint(p)

				e, err := db.ExchangeByName("IDEX")
				if err != nil {
					log.Panicf("Error fetching IDEX exchange: %v\n", err)
				}

				m, err := db.MarketByExTokAddr(e.ID, tknAddr)
				if err != nil {
					continue // skip if market not there
				}

				o := &orderbook.Order{
					MarketID:    m.ID,
					Volume:      vol,
					Price:       price,
					IsBuy:       isBuy,
					ExpireBlock: strconv.Itoa(r.OrderInserted.Expires),
					UserAddress: r.OrderInserted.User,
					Nonce:       strconv.Itoa(r.OrderInserted.Nonce),
					Hash:        r.OrderInserted.Hash,
				}

				if err := db.InsertOrder(o); err != nil {
					log.Panicf("Error inserting order %+v: %v\n", r.OrderInserted, err)
				}
			}
			if r.Error != nil {
				fmt.Printf("got an error: %+v\n", r.Error)
			}
		case <-ctx.Done():
			err := i.Socket.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error writing close message:", err)
				return
			}

			<-time.After(time.Second)
			return
		}
	}
}

func loadConfig(fileName string) {
	c, err := config.Read(fileName)
	if err != nil {
		log.Fatal(err)
	}
	Config = c
}

// RedisKey formats the values into a consistent key format
func RedisKey(exchange, token, orderID string) string {
	return fmt.Sprintf("%v-%v-%v", exchange, token, orderID)
}
