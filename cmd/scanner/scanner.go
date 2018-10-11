package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/mathieugilbert/aabot/cmd/config"
	"github.com/mathieugilbert/aabot/orderbook"

	_ "github.com/jinzhu/gorm/dialects/postgres" // db driver
)

// Opportunity -
type Opportunity struct {
	Bid    *orderbook.Order
	Ask    *orderbook.Order
	Profit *big.Float
	Volume *big.Float
}

// String override stringer output
func (o Opportunity) String() string {
	pct := new(big.Float).Quo(o.Profit, o.Volume)
	pct.Mul(pct, big.NewFloat(100))

	msg := fmt.Sprintf("Buy on market: %v, price: %v.", o.Ask.MarketID, o.Ask.Price)
	msg += fmt.Sprintf("Sell on market: %v, price: %v.", o.Bid.MarketID, o.Bid.Price)
	msg += fmt.Sprintf("Volume: %v, Profit: %v (%v%%)\n", o.Volume, o.Profit, pct)
	return msg
}

// DB -
type DB struct {
	*orderbook.DB
}

var (
	// Config -
	Config *config.Configuration
	// ConfigFile -
	ConfigFile = "../../private/mainnet/config.json"
)

func main() {
	loadConfig(ConfigFile)

	odb, err := orderbook.NewDB(Config.DBString())
	if err != nil {
		log.Fatal(err)
	}
	odb.LogMode(false)

	db := &DB{odb}

	opp := make(chan Opportunity)

	log.Println("Scanning for arbitrage opportunities:")

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Printf(".")
				db.ScanPrices(opp)
			}
		}
	}()

	for {
		select {
		case o := <-opp:
			log.Printf("\nFound arbitrage opportunity! %s\n", o)
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

// ScanPrices looks for open orders with bid > ask price
// Bid is someone's buy order, which I can sell to.
// Ask is someone's sell order, which I can buy from.
func (db *DB) ScanPrices(opp chan<- Opportunity) {
	ts, err := db.Tokens()
	if err != nil {
		log.Panic("Error fetching tokens")
	}

	for _, t := range ts {
		highBid, err := db.HighestBid(t.ID)
		if err != nil {
			log.Panicf("Error getting HightestBid(%v): %v\n", t.ID, err)
		}

		lowAsk, err := db.LowestAsk(t.ID)
		if err != nil {
			log.Panicf("Error getting LowestAsk(%v): %v\n", t.ID, err)
		}

		if len(highBid) == 0 || len(lowAsk) == 0 {
			continue
		}

		bp := new(big.Float)
		if _, ok := bp.SetString(highBid[0].Price); !ok {
			log.Panicf("Error converting price to big.Float: %v, %v\n", highBid[0].Price, err)
		}
		ap := new(big.Float)
		if _, ok := ap.SetString(lowAsk[0].Price); !ok {
			log.Panicf("Error converting price to big.Float: %v, %v\n", lowAsk[0].Price, err)
		}

		// quick check to ensure bid > ask
		if bp.Cmp(ap) <= 0 {
			continue
		}

		// get the minimum volume of the two
		bvol := new(big.Float)
		if _, ok := bvol.SetString(highBid[0].Volume); !ok {
			log.Panicf("Invalid volume: %v\n", highBid[0].Volume)
		}
		bv := new(big.Float).Quo(bvol, big.NewFloat(math.Pow10(18)))
		// convert to ETH
		bv.Mul(bv, bp)

		avol := new(big.Float)
		if _, ok := avol.SetString(lowAsk[0].Volume); !ok {
			log.Panicf("Invalid volume: %v\n", lowAsk[0].Volume)
		}
		av := new(big.Float).Quo(avol, big.NewFloat(math.Pow10(18)))
		// convert to ETH
		av.Mul(av, ap)

		var vol *big.Float // in ETH
		if bv.Cmp(av) > 0 {
			vol = av
		} else {
			vol = bv
		}

		// look up each side's exchange to get the fee
		be, err := db.ExchangeByMarket(highBid[0].MarketID)
		if err != nil {
			log.Panicf("Error getting ExchangeByMarket(%v): %v\n", highBid[0].MarketID, err)
		}

		ae, err := db.ExchangeByMarket(lowAsk[0].MarketID)
		if err != nil {
			log.Panicf("Error getting ExchangeByMarket(%v): %v\n", lowAsk[0].MarketID, err)
		}

		profit := new(big.Float)
		profit.Mul(vol, profit.Sub(profit.Mul(profit.Mul(profit.Quo(bp, ap), big.NewFloat(1-be.Fee)), big.NewFloat(1-ae.Fee)), big.NewFloat(1)))

		gasFeeDoubled := big.NewFloat(0.00011 * 2)

		if profit.Cmp(gasFeeDoubled) > 0 {
			opp <- Opportunity{
				Bid:    highBid[0],
				Ask:    lowAsk[0],
				Profit: profit,
				Volume: vol,
			}
		}
	}

}

// HighestBid returns top count highest price buy orders for token
func (db *DB) HighestBid(tid int) (os []*orderbook.Order, err error) {
	//err = db.Table("orders").Where("market_id = ? AND is_buy = true", mid).Order("price desc").Limit(1).Find(&os).Error
	err = db.Joins("left join markets on markets.id = orders.market_id").Where("markets.token_id = ? AND orders.is_buy = true", tid).Order("price desc").Limit(1).Find(&os).Error
	return
}

// LowestAsk returns bottom count lowest price sell orders for token
func (db *DB) LowestAsk(tid int) (os []*orderbook.Order, err error) {
	//err = db.Table("orders").Where("market_id = ? AND is_buy = false", mid).Order("price asc").Limit(1).Find(&os).Error
	err = db.Joins("left join markets on markets.id = orders.market_id").Where("markets.token_id = ? AND orders.is_buy = false", tid).Order("price asc").Limit(1).Find(&os).Error
	return
}
