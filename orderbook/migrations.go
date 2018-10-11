package orderbook

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

// Migrate models with gorm
func Migrate(s string) {
	db, err := gorm.Open("postgres", s)
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %v", err))
	}
	if err = db.DB().Ping(); err != nil {
		panic(fmt.Errorf("unable to ping database: %v", err))
	}
	db.LogMode(true)
	defer db.Close()

	if err = start(db); err != nil {
		panic(fmt.Errorf("Could not migrate: %v", err))
	}
	fmt.Println("Migration ran successfully")
}

// IDs are timestamps from command line:
// > date +"%Y%m%d%H%M%S"
func start(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create tables
		{
			ID: "20180909110929",
			Migrate: func(tx *gorm.DB) error {
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
				type Token struct {
					ID      int    `gorm:"primary_key"`
					Name    string `gorm:"not null;index"`
					Address string `gorm:"not null"`
				}
				type Exchange struct {
					ID      int     `gorm:"primary_key"`
					Name    string  `gorm:"not null;index"`
					Address string  `gorm:"not null"`
					Fee     float64 `gorm:"not null"`
				}
				type Market struct {
					ID         int    `gorm:"primary_key"`
					ExchangeID int    `gorm:"not null"`
					TokenID    int    `gorm:"not null"`
					Balance    string `gorm:"default:'0'"`
					Active     bool   `gorm:"not null"`
				}
				type Wallet struct {
					ID      int    `gorm:"primary_key"`
					TokenID int    `gorm:"token_id"`
					Balance string `gorm:"balance"`
				}

				if err := tx.CreateTable(&Token{}).Error; err != nil {
					return err
				}

				if err := tx.CreateTable(&Exchange{}).Error; err != nil {
					return err
				}

				if err := tx.CreateTable(&Market{}).Error; err != nil {
					return err
				}
				if err := tx.Model(&Market{}).AddForeignKey("exchange_id", "exchanges(id)", "RESTRICT", "RESTRICT").Error; err != nil {
					return err
				}
				if err := tx.Model(&Market{}).AddForeignKey("token_id", "tokens(id)", "RESTRICT", "RESTRICT").Error; err != nil {
					return err
				}

				if err := tx.CreateTable(&Order{}).Error; err != nil {
					return err
				}
				if err := tx.Model(&Order{}).AddForeignKey("market_id", "markets(id)", "RESTRICT", "RESTRICT").Error; err != nil {
					return err
				}

				if err := tx.CreateTable(&Wallet{}).Error; err != nil {
					return err
				}
				if err := tx.Model(&Wallet{}).AddForeignKey("token_id", "tokens(id)", "RESTRICT", "RESTRICT").Error; err != nil {
					return err
				}

				return nil
			},
		},
	})

	return m.Migrate()
}
