package orderbook

import (
	"strings"
	"time"
)

// Order holds standing orders from exchanges, with enough info to take the trade
type Order struct {
	ID          int       `gorm:"primary_key"`
	CreatedAt   time.Time `gorm:"not null"`
	MarketID    int       `gorm:"not null"`       // FK
	ExchangeOID string    `gorm:"unique"`         // Exchange's ID for the Order
	Volume      string    `gorm:"not null"`       // in token
	Price       string    `gorm:"not null;index"` // in ETH
	IsBuy       bool      `gorm:"not null"`       // buy/sell
	ExpireBlock string    ``                      // block number order will expire after
	UserAddress string    ``                      // address of user who placed order
	Nonce       string    ``                      // hash var
	V           int       ``                      // hash var
	R           string    ``                      // hash var
	S           string    ``                      // hash var
}

// CreateOrder inserts the order
func (db *DB) CreateOrder(o *Order) error {
	return db.Create(o).Error
}

// BulkInsertOrders creates all the orders as bulk insert
func (db *DB) BulkInsertOrders(orders []*Order) error {
	cmd := "INSERT INTO orders(created_at, market_id, exchange_o_id, volume, price, is_buy, expire_block, user_address, nonce, v, r, s) VALUES "
	const slots = "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	var inserts []string
	vals := []interface{}{}
	t := time.Now()

	for _, o := range orders {
		inserts = append(inserts, slots)
		vals = append(vals, t, o.MarketID, o.ExchangeOID, o.Volume, o.Price, o.IsBuy, o.ExpireBlock, o.UserAddress, o.Nonce, o.V, o.R, o.S)
	}

	cmd = cmd + strings.Join(inserts, ",")

	return db.Exec(cmd, vals...).Error
}

// HighestBuys returns top count highest price buy orders for token
func (db *DB) HighestBuys(mid int, count int) (os []*Order, err error) {
	err = db.Where(&Order{MarketID: mid, IsBuy: true}).Order("price desc").Limit(count).Find(&os).Error
	return
}

// LowestSells returns bottom count lowest price sell orders for token
func (db *DB) LowestSells(tid int, count int) (os []*Order, err error) {
	// gorm ignores "false" value, so doesn't work as .Where(struct) like for the buys
	err = db.Table("orders").Where("token_id = ? AND is_buy = false", tid).Order("price asc").Limit(count).Find(&os).Error
	return
}
