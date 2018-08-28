package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"

	"github.com/go-redis/redis"
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

	// going to wait for all exchanges to load
	var wg sync.WaitGroup
	wg.Add(len(es))

	for _, e := range es {
		go func(e *orderbook.Exchange, wg *sync.WaitGroup) {
			e.LoadOrderbooks(db, &orderbook.ExchangeServices{GoEd: ed})
			wg.Done()
		}(e, &wg)
	}

	wg.Wait()

	// combine this into the LoadOrderbooks?
	// pass a single interrupt into each loop?
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for _, e := range es {
		go func(e *orderbook.Exchange) {
			e.SyncOrderbook(db, &orderbook.ExchangeServices{GoEd: ed}, interrupt)
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
