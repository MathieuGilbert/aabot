package orderbook

import (
	"log"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/mathieugilbert/aabot/adapters"
	goed "github.com/mathieugilbert/go-etherdelta"
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

// MarketByExTok fetches by exchange id and token id
func (db *DB) MarketByExTok(eid, tid int) (*Market, error) {
	m := &Market{}
	if err := db.Where(&Market{ExchangeID: eid, TokenID: tid}).First(m).Error; err != nil {
		return nil, err
	}
	return m, nil
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

// can this be done with the monitoring websocket connection?
// 2018/08/27 23:11:03 Error getting ED order book for market 18: Error getting orders, got null
// 2018/08/27 23:11:03 New trades, refreshed orderbook for token KIN on exchange EtherDelta

func (m *Market) LoadOrderbook(db *DB, s *ExchangeServices, wg *sync.WaitGroup) {
	var volumeThreshold = 0.001 * math.Pow10(18)

	if m.Token.Name == "ETH" || !m.Active {
		wg.Done()
		return
	}

	opts := &goed.GetOrderBookOpts{
		TokenAddress: m.Token.Address,
	}

	orders, err := s.GoEd.GetOrderBook(opts)
	if err != nil {
		log.Printf("Error getting ED order book for market %v: %v\n", m.ID, err)
		wg.Done()
		return
	}

	var os []*Order

	for _, order := range orders.Buys {
		vol, err := strconv.ParseFloat(order.AvailableVolumeBase, 32)
		if err != nil {
			log.Printf("ParseFloat(%v) error: %v\n", order.AvailableVolumeBase, err)
			continue
		}
		if vol < volumeThreshold {
			continue
		}

		o := &Order{
			MarketID:    m.ID,
			ExchangeOID: order.Id,
			Volume:      order.AvailableVolume,
			Price:       order.Price,
			IsBuy:       true,
			ExpireBlock: order.Expires,
			UserAddress: order.User,
			Nonce:       order.Nonce,
			V:           order.V,
			R:           order.R,
			S:           order.S,
		}

		os = append(os, o)
	}

	for _, order := range orders.Sells {
		vol, err := strconv.ParseFloat(order.AvailableVolumeBase, 32)
		if err != nil {
			log.Printf("ParseFloat(%v) error: %v\n", order.AvailableVolumeBase, err)
			continue
		}
		if vol < volumeThreshold {
			continue
		}

		o := &Order{
			MarketID:    m.ID,
			ExchangeOID: order.Id,
			Volume:      order.AvailableVolume,
			Price:       order.Price,
			IsBuy:       false,
			ExpireBlock: order.Expires,
			UserAddress: order.User,
			Nonce:       order.Nonce,
			V:           order.V,
			R:           order.R,
			S:           order.S,
		}

		os = append(os, o)
	}

	if err := db.BulkInsertOrders(os); err != nil {
		log.Printf("Error inserting orders for market %v: %v\n", m.ID, err)
		wg.Done()
		return
	}

	wg.Done()
}
