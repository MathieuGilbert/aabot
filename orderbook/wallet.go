package orderbook

import (
	"log"
	"time"

	"github.com/mathieugilbert/aabot/adapters"
)

// Wallet holds token balances
type Wallet struct {
	ID      int    `gorm:"primary_key" json:"id"`
	TokenID int    `gorm:"token_id" json:"token_id"`
	Balance string `gorm:"balance" json:"balance"`
}

// Balance returns the balance of the token
func (db *DB) Balance(tid int) (string, error) {
	w := &Wallet{}
	err := db.Where(&Wallet{TokenID: tid}).First(w).Error
	if err != nil {
		return "", err
	}

	return w.Balance, nil
}

// Wallets returns all the wallets
func (db *DB) Wallets() (ws []*Wallet, err error) {
	err = db.Order("token_id").Find(&ws).Error
	return
}

// setWalletBalance updates the wallet balance
func (db *DB) setWalletBalance(id int, bal string) error {
	w := &Wallet{ID: id}
	return db.Model(w).Update("balance", bal).Error
}

func (db *DB) setMarketBalance(id int, bal string) error {
	m := &Market{ID: id}
	return db.Model(m).Update("balance", bal).Error
}

// SyncBalances all token balances in wallet
func (db *DB) SyncBalances(eth *adapters.Ethereum, addr string) error {
	ts, err := db.Tokens()
	if err != nil {
		log.Printf("Error fetching tokens: %v\n", err)
		return err
	}

	for _, t := range ts {
		ticker := time.NewTicker(10 * time.Second)

		go func(t *Token) {
			for {
				<-ticker.C

				bal := eth.TokenBalance(addr, t.Name, t.Address)
				if err := db.setWalletBalance(t.ID, bal); err != nil {
					log.Printf("Error updating balance: %v\n", err)
				}
			}
		}(t)
	}

	return nil
}
