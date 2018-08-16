package orderbook

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
)

// DB wraps gorm.DB
type DB struct {
	*gorm.DB
}

// Datastore implements available DB methods
type Datastore interface {
	Seed(string) error
	ClearTempTables()
	Exchanges() ([]*Exchange, error)
	ExchangeByName(string) (Exchange, error)
	Tokens() ([]*Token, error)
	Token(int) (*Token, error)
	TokenByName(string) (Token, error)
	Markets() ([]*Market, error)
	MarketByExTok(int, int) (*Market, error)
	CreateOrder(*Order) error
	BulkInsertOrders([]*Order) error
	HighestBuys(int, int) ([]*Order, error)
	LowestSells(int, int) ([]*Order, error)
}

// NewDB creates a new connection
func NewDB(source string) (*DB, error) {
	db, err := gorm.Open("postgres", source)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// Seed the database from provided file
func (db *DB) Seed(fileName string) error {
	type DisabledToken struct {
		Exchange string
		Token    string
	}
	type Seed struct {
		Exchanges      []Exchange
		Tokens         []Token
		DisabledTokens []DisabledToken
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	s := &Seed{}
	if err = json.Unmarshal(file, &s); err != nil {
		log.Panic(err)
		return err
	}

	for _, e := range s.Exchanges {
		if err := db.FirstOrCreate(&Exchange{}, e).Error; err != nil {
			return err
		}
	}

	for _, t := range s.Tokens {
		if err := db.FirstOrCreate(&Token{}, t).Error; err != nil {
			return err
		}
	}

	// assume all enabled unless included
	for _, e := range s.Exchanges {
		for _, t := range s.Tokens {
			// make all permutations active
			active := true

			// check if included
			for _, dt := range s.DisabledTokens {
				if dt.Exchange == e.Name && dt.Token == t.Name {
					active = false
				}
			}

			ti, _ := db.TokenByName(t.Name)
			ei, _ := db.ExchangeByName(e.Name)

			rec := &Market{
				ExchangeID: ei.ID,
				TokenID:    ti.ID,
				Active:     active,
			}

			if err := db.FirstOrCreate(&Market{}, rec).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// ClearTempTables empties the Orders, Tokens, Exchanges
func (db *DB) ClearTempTables() {
	db.Exec("DELETE FROM orders;")
	db.Exec("DELETE FROM tokens;")
	db.Exec("DELETE FROM exchanges;")
	db.Exec("DELETE FROM markets;")
}
