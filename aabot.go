package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	_ "github.com/jinzhu/gorm/dialects/postgres" // db driver
	"github.com/mathieugilbert/aabot/adapters"
	"github.com/mathieugilbert/aabot/cmd/config"

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
	//db.Seed(SeedFile)

	// Ethereum network
	eth, err := adapters.NewEthereum(Config.Ethereum.ProviderURI)
	if err != nil {
		log.Panicf("Couldn't connect to Ethereum network: %v\n%v\n", Config.Ethereum.ProviderURI, err)
	}

	e, err := db.ExchangeByName("EtherDelta")
	if err != nil {
		log.Panicf("Couldn't find EtherDelta in DB: %v\n", err)
	}
	log.Printf("exchange id: %v", e.ID)

	// EtherDelta websocket service
	ed := goed.NewForkDelta(&goed.Options{
		ProviderURI: e.WebsocketURI,
	})

	log.Println("got new ED WS service")

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
		GoEd:     ed,
		Redis:    r,
		DB:       db,
	}

	if err := db.SyncBalances(eth, Config.Ethereum.UserAddress); err != nil {
		log.Panicf("Unable to sync wallet balances: %v\n", err)
	}
	if err := db.SyncMarkets(eth, Config.Ethereum.UserAddress); err != nil {
		log.Panicf("Unable to sync exchange balances: %v\n", err)
	}

	es, err := db.Exchanges()
	if err != nil {
		log.Panicf("Unable to retrieve exchanges: %v\n", err)
	}

	for _, e := range es {
		go func(e *orderbook.Exchange) {
			e.LoadOrderbooks(db, &orderbook.ExchangeServices{GoEd: ed})
		}(e)
	}

	web.Start(ss)
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

func run() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	//u := url.URL{Scheme: "ws", Host: "https://api.forkdelta.com", Path: "/echo"}
	//log.Printf("connecting to %s", u.String())

	log.Println("Connecting to EtherDelta ws...")

	c, _, err := websocket.DefaultDialer.Dial("wss://socket.etherdelta.com/socket.io/?transport=websocket", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	log.Println("Connected to EtherDelta ws.")

	err = c.WriteMessage(websocket.TextMessage, []byte("42['getMarket', {'token':''}]"))
	if err != nil {
		log.Println("error: ", err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("ReadMessage error:", err)
				return
			}
			log.Printf("received: %s", message)
		}
	}()

	err = c.WriteMessage(websocket.TextMessage, []byte("~~~~~~~~~~~ AWFUL ~~~~~~~~~~~~"))
	if err != nil {
		log.Println("WriteMessage error:", err)
		return
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("WriteMessage error:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
