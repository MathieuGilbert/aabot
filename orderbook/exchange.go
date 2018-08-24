package orderbook

import (
	"log"
	"math"
	"strconv"

	goed "github.com/mathieugilbert/go-etherdelta"
)

// Exchange holds exchange-specific details
type Exchange struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `gorm:"not null;index" json:"name"`
	Address      string    `gorm:"not null" json:"address"`
	WebsocketURI string    `gorm:"not null" json:"-"`
	Fee          float32   `gorm:"not null" json:"fee"`
	Markets      []*Market `json:"markets"`
}

// ExchangeServices holds instantiated services for each exchange
type ExchangeServices struct {
	GoEd *goed.Service
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

// ExchangeByName returns an exchange by name
func (db *DB) ExchangeByName(name string) (e Exchange, err error) {
	err = db.Where(&Exchange{Name: name}).First(&e).Error
	return
}

// LoadOrderbooks fetches and stores open orders for active markets on the exchange
func (e *Exchange) LoadOrderbooks(db *DB, s *ExchangeServices) error {
	ms, err := db.ExchangeMarketsJoined(e.ID)
	if err != nil {
		log.Printf("Error getting markets for exchange %v: %v", e.Name, err)
		return err
	}

	for _, m := range ms {
		if m.Token.Name == "ETH" || !m.Active {
			continue
		}

		switch e.Name {
		case "EtherDelta":
			opts := &goed.GetOrderBookOpts{
				TokenAddress: m.Token.Address,
			}

			orders, err := s.GoEd.GetOrderBook(opts)
			if err != nil {
				log.Printf("Error getting ED order book: %v", err)
				return err
			}

			var os []*Order

			for _, order := range orders.Buys {
				// skip if volume ~ gas fee
				vol, err := strconv.ParseFloat(order.AvailableVolumeBase, 32)
				if err != nil || vol < 0.0001*math.Pow10(18) {
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
				// skip if volume ~ gas fee
				vol, err := strconv.ParseFloat(order.AvailableVolumeBase, 32)
				if err != nil || vol < 0.0001*math.Pow10(18) {
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
				log.Printf("Error inserting orders for exchange %v, market %v: %v\n", e.ID, m.ID, err)
				return err
			}
		}
	}

	return nil
}
