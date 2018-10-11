package orderbook

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mathieugilbert/aabot/adapters"
	goed "github.com/mathieugilbert/go-etherdelta"
	idex "github.com/mathieugilbert/go-idex"
)

// Market holds information about tokens per exchange
type Market struct {
	ID         int      `gorm:"primary_key" json:"id"`
	ExchangeID int      `gorm:"not null" json:"exchange_id"`
	TokenID    int      `gorm:"not null" json:"token_id"`
	Balance    string   `gorm:"default:'0'" json:"balance"`
	Active     bool     `gorm:"not null" json:"active"`
	Exchange   Exchange `json:"exchange"`
	Token      Token    `json:"token"`
	Orders     []*Order `json:"orders"`
}

// Markets returns all the markets
func (db *DB) Markets() (ms []*Market, err error) {
	err = db.Order("token_id").Find(&ms).Error
	return
}

// MarketsJoined returns all the markets with Exchange and Token filled
func (db *DB) MarketsJoined() (ms []*Market, err error) {
	err = db.Preload("Token").Preload("Exchange").Order("token_id").Find(&ms).Error
	if err != nil {
		return nil, err
	}

	for _, m := range ms {
		os, err := db.PreviewOrderbook(m.ID, 5)
		if err != nil {
			return nil, err
		}
		m.Orders = os
	}

	return
}

// Market returns the token with the id
func (db *DB) Market(id int) (*Market, error) {
	m := &Market{}
	err := db.First(m, id).Error
	return m, err
}

// MarketByExTokAddr fetches by exchange id and token address
func (db *DB) MarketByExTokAddr(eid int, tokenAddr string) (m *Market, err error) {
	m = &Market{}
	err = db.Joins("left join tokens on tokens.id = markets.token_id").Where("tokens.address = UPPER(?) AND markets.exchange_id = ?", tokenAddr, eid).First(m).Error
	return
}

// ExchangeMarketsJoined returns all markets for the exchange, preloading token and exchange
func (db *DB) ExchangeMarketsJoined(eid int) (ms []*Market, err error) {
	err = db.Preload("Token").Preload("Exchange").Order("token_id").Where(&Market{ExchangeID: eid}).Find(&ms).Error
	return
}

// SyncMarketBalances store token balances from all market exchanges
func (db *DB) SyncMarketBalances(eth *adapters.Ethereum, addr string) error {
	ms, err := db.MarketsJoined()
	if err != nil {
		log.Printf("Error fetching markets: %v\n", err)
		return err
	}

	for _, m := range ms {
		ticker := time.NewTicker(10 * time.Second)

		go func(m *Market) {
			for {
				<-ticker.C

				bal := eth.ExchangeTokenBalance(m.Exchange.Name, m.Exchange.Address, m.Token.Address, addr)
				if err := db.setMarketBalance(m.ID, bal); err != nil {
					log.Printf("Error updating balance: %v\n", err)
				}
			}
		}(m)
	}

	return nil
}

// LoadOrderbookEtherDelta loads orderbook for the market on EtherDelta
// 2018/08/27 23:11:03 Error getting ED order book for market 18: Error getting orders, got null
// 2018/08/27 23:11:03 New trades, refreshed orderbook for token KIN on exchange EtherDelta
func (m *Market) LoadOrderbookEtherDelta(db *DB, wg *sync.WaitGroup) {
	if m.Token.Name == "ETH" || !m.Active {
		wg.Done()
		return
	}

	ed := goed.NewForkDelta(&goed.Options{})
	ob, err := ed.GetOrderBook(&goed.GetOrderBookOpts{TokenAddress: m.Token.Address})
	if err != nil {
		log.Printf("Error getting ED order book for market %v: %v\n", m.ID, err)
		wg.Done()
		return
	}

	os, err := m.Exchange.OrderbookToOrdersEtherDelta(db, ob)
	if err != nil {
		log.Printf("Error processing new EtherDelta orders for market %v: %v\n", m.ID, err)
		wg.Done()
		return
	}

	if len(os) == 0 {
		wg.Done()
		return
	}

	if err := db.BulkInsertOrders(os); err != nil {
		log.Printf("Error inserting orders for market %v: %v\n", m.ID, err)
		wg.Done()
		return
	}

	wg.Done()
}

// LoadOrderbookIDEX loads orderbook for the market
func (m *Market) LoadOrderbookIDEX(db *DB, wg *sync.WaitGroup) {
	if m.Token.Name == "ETH" || !m.Active {
		wg.Done()
		return
	}

	i := idex.New()
	ob, err := i.API.OrderBook(fmt.Sprintf("ETH_%v", m.Token.Name))
	if err != nil {
		log.Printf("Error getting open orders for market %v: %v\n", m.ID, err)
		wg.Done()
		return
	}

	os, err := m.Exchange.OrderbookToOrdersIdex(db, ob)
	if err != nil {
		log.Printf("Error processing new Idex orders for market %v: %v\n", m.ID, err)
		wg.Done()
		return
	}

	if len(os) == 0 {
		wg.Done()
		return
	}

	if err := db.BulkInsertOrders(os); err != nil {
		log.Printf("Error inserting orders for market %+v: %v\n", m, err)
		wg.Done()
		return
	}

	wg.Done()
}
