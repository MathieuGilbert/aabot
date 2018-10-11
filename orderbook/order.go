package orderbook

import (
	"log"
	"sort"
	"strings"
	"time"
)

// Order holds standing orders from exchanges, with enough info to take the trade
type Order struct {
	ID          int       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	MarketID    int       `gorm:"not null" json:"market_id"`                      // FK
	ExchangeOID string    `json:"exchange_o_id"`                                  // Exchange's ID for the Order
	Volume      string    `gorm:"not null" sql:"type:numeric" json:"volume"`      // in token
	Price       string    `gorm:"not null;index" sql:"type:numeric" json:"price"` // in ETH
	IsBuy       bool      `gorm:"not null" json:"is_buy"`                         // buy/sell
	ExpireBlock string    `json:"expire_block"`                                   // block number order will expire after
	UserAddress string    `json:"user_address"`                                   // address of user who placed order
	Nonce       string    `json:"nonce"`                                          // hash var
	V           int       `json:"v"`                                              // hash var
	R           string    `json:"r"`                                              // hash var
	S           string    `json:"s"`                                              // hash var
	Hash        string    `gorm:"index" json:"hash"`                              // computed hash
}

// BulkInsertOrders creates all the orders as bulk insert
func (db *DB) BulkInsertOrders(orders []*Order) error {
	db.LogMode(false)

	cmd := "INSERT INTO orders(created_at, market_id, exchange_o_id, volume, price, is_buy, expire_block, user_address, nonce, v, r, s, hash) VALUES "
	const slots = "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	var inserts []string
	vals := []interface{}{}
	t := time.Now()

	for _, o := range orders {
		inserts = append(inserts, slots)
		vals = append(vals, t, o.MarketID, o.ExchangeOID, o.Volume, o.Price, o.IsBuy, o.ExpireBlock, o.UserAddress, o.Nonce, o.V, o.R, o.S, o.Hash)
	}

	cmd = cmd + strings.Join(inserts, ",")

	err := db.Exec(cmd, vals...).Error
	db.LogMode(true)

	if err != nil {
		log.Println(err)
	}

	return err
}

// PreviewOrderbook returns count buys and sells nearest midpoint for a market
func (db *DB) PreviewOrderbook(mid int, count int) ([]*Order, error) {
	buys, err := db.HighestBuys(mid, count)
	if err != nil {
		return nil, err
	}

	sells, err := db.LowestSells(mid, count)
	if err != nil {
		return nil, err
	}

	sort.Slice(sells, func(i int, j int) bool {
		return sells[i].Price > sells[j].Price
	})

	return append(sells, buys...), nil
}

// InsertOrder creates a new order
func (db *DB) InsertOrder(o *Order) error {
	return db.Create(o).Error
}

// HighestBuys returns top count highest price buy orders for token
func (db *DB) HighestBuys(mid int, count int) (os []*Order, err error) {
	err = db.Table("orders").Where("market_id = ? AND is_buy = true", mid).Order("price desc").Limit(count).Find(&os).Error
	return
}

// LowestSells returns bottom count lowest price sell orders for token
func (db *DB) LowestSells(mid int, count int) (os []*Order, err error) {
	err = db.Table("orders").Where("market_id = ? AND is_buy = false", mid).Order("price asc").Limit(count).Find(&os).Error
	return
}

// DeleteOrderByHash deletes the order with the hash
func (db *DB) DeleteOrderByHash(hash string) (err error) {
	err = db.Where("hash = ?", hash).Delete(&Order{}).Error
	return
}
