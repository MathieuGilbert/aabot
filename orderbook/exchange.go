package orderbook

import (
	"log"
	"math"
	"strconv"
	"sync"

	goed "github.com/mathieugilbert/go-etherdelta"
	idex "github.com/mathieugilbert/go-idex"
)

// Exchange holds exchange-specific details
type Exchange struct {
	ID      int       `gorm:"primary_key" json:"id"`
	Name    string    `gorm:"not null;index" json:"name"`
	Address string    `gorm:"not null" json:"address"`
	Fee     float64   `gorm:"not null" json:"fee"`
	Markets []*Market `json:"markets"`
}

// Exchanges returns all the exchanges
func (db *DB) Exchanges() (es []*Exchange, err error) {
	err = db.Find(&es).Error
	return
}

// ExchangesJoined returns all the exchanges with their markets
func (db *DB) ExchangesJoined() (es []*Exchange, err error) {
	err = db.Preload("Markets").Find(&es).Order("id").Error
	if err != nil {
		return nil, err
	}

	for _, e := range es {
		for _, m := range e.Markets {
			os, err := db.PreviewOrderbook(m.ID, 5)
			if err != nil {
				return nil, err
			}
			m.Orders = os
		}
	}

	return
}

// ExchangeByMarket returns the Exchange for the market ID
func (db *DB) ExchangeByMarket(mid int) (e Exchange, err error) {
	err = db.Joins("left join markets on markets.exchange_id = exchanges.id").Where("markets.id = ?", mid).First(&e).Error
	return
}

// ExchangeByName returns an exchange by name
func (db *DB) ExchangeByName(name string) (e Exchange, err error) {
	err = db.Where(&Exchange{Name: name}).First(&e).Error
	return
}

// LoadOrderbooks fetches and stores open orders for active markets on the exchange
func (e *Exchange) LoadOrderbooks(db *DB) error {
	ms, err := db.ExchangeMarketsJoined(e.ID)
	if err != nil {
		log.Printf("Error getting markets for exchange %v: %v\n", e.Name, err)
		return err
	}

	// going to wait until all markets finish
	var wg sync.WaitGroup
	wg.Add(len(ms))

	for _, m := range ms {
		switch e.Name {
		case "EtherDelta":
			go m.LoadOrderbookEtherDelta(db, &wg)
		case "IDEX":
			go m.LoadOrderbookIDEX(db, &wg)
		}
	}

	wg.Wait()
	return nil
}

// ProcessNewOrdersEtherDelta processes orders in the orderbook
func (e *Exchange) ProcessNewOrdersEtherDelta(db *DB, ob *goed.OrderBook) (int, error) {
	os, err := e.OrderbookToOrdersEtherDelta(db, ob)
	if err != nil {
		return 0, err
	}
	if len(os) == 0 {
		return 0, nil
	}

	if err := db.BulkInsertOrders(os); err != nil {
		return 0, err
	}
	return len(os), nil
}

// OrderbookToOrdersEtherDelta converts an ED orderbook into slice of aabot orders
func (e *Exchange) OrderbookToOrdersEtherDelta(db *DB, ob *goed.OrderBook) ([]*Order, error) {
	var volumeThreshold = 0.001 * math.Pow10(18) // wei
	var os []*Order
	mids := make(map[string]int) // token address to market id

	for _, order := range ob.Buys {
		vol, err := strconv.ParseFloat(order.AvailableVolumeBase, 32)
		if err != nil {
			return nil, err
		}
		if vol < volumeThreshold {
			// skip order since too small
			continue
		}

		// cache this in a map
		key := order.TokenGet
		if mid, ok := mids[key]; ok {
			if mid == 0 {
				continue
			}
		} else {
			m, err := db.MarketByExTokAddr(e.ID, key)
			if err != nil {
				// if not in db, we don't care about it
				mids[key] = 0
				continue
			}
			mids[key] = m.ID
		}

		o := &Order{
			MarketID:    mids[key],
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

	for _, order := range ob.Sells {
		vol, err := strconv.ParseFloat(order.AvailableVolumeBase, 32)
		if err != nil {
			return nil, err
		}
		if vol < volumeThreshold {
			// skip order since too small
			continue
		}

		// cache this in a map
		key := order.TokenGive
		if mid, ok := mids[key]; ok {
			if mid == 0 {
				continue
			}
		} else {
			m, err := db.MarketByExTokAddr(e.ID, key)
			if err != nil {
				// if not in db, we don't care about it
				mids[key] = 0
				continue
			}
			mids[key] = m.ID
		}

		o := &Order{
			MarketID:    mids[key],
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

	return os, nil
}

// OrderbookToOrdersIdex converts an IDEX orderbook into slice of aabot orders
func (e *Exchange) OrderbookToOrdersIdex(db *DB, ob *idex.OrderBook) ([]*Order, error) {
	var volumeThreshold = 0.001 // ether
	var os []*Order
	mids := make(map[string]int) // token address to market id

	for _, order := range ob.Bids {
		vol, err := strconv.ParseFloat(order.Total, 32)
		if err != nil {
			return nil, err
		}

		if vol < volumeThreshold {
			// skip order since too small
			continue
		}

		// cache this in a map
		key := order.Params.TokenBuy
		if mid, ok := mids[key]; ok {
			if mid == 0 {
				continue
			}
		} else {
			m, err := db.MarketByExTokAddr(e.ID, key)
			if err != nil {
				// if not in db, we don't care about it
				mids[key] = 0
				continue
			}
			mids[key] = m.ID
		}

		o := &Order{
			MarketID:    mids[key],
			Volume:      order.Params.AmountBuy,
			Price:       order.Price,
			IsBuy:       true,
			ExpireBlock: strconv.Itoa(order.Params.Expires),
			UserAddress: order.Params.User,
			Nonce:       strconv.Itoa(order.Params.Nonce),
			Hash:        order.OrderHash,
		}

		os = append(os, o)
	}

	for _, order := range ob.Asks {
		vol, err := strconv.ParseFloat(order.Total, 32)
		if err != nil {
			return nil, err
		}
		if vol < volumeThreshold {
			// skip order since too small
			continue
		}

		// cache this in a map
		key := order.Params.TokenSell
		if mid, ok := mids[key]; ok {
			if mid == 0 {
				continue
			}
		} else {
			m, err := db.MarketByExTokAddr(e.ID, key)
			if err != nil {
				// if not in db, we don't care about it
				mids[key] = 0
				continue
			}
			mids[key] = m.ID
		}

		o := &Order{
			MarketID:    mids[key],
			Volume:      order.Params.AmountSell,
			Price:       order.Price,
			IsBuy:       false,
			ExpireBlock: strconv.Itoa(order.Params.Expires),
			UserAddress: order.Params.User,
			Nonce:       strconv.Itoa(order.Params.Nonce),
			Hash:        order.OrderHash,
		}

		os = append(os, o)
	}

	return os, nil
}
