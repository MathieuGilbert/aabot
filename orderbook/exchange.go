package orderbook

// Exchange holds exchange-specific details
type Exchange struct {
	ID           int    `gorm:"primary_key"`
	Name         string `gorm:"not null;index"`
	Address      string `gorm:"not null"`
	WebsocketURI string `gorm:"not null"`
}

// Exchanges returns all the exchanges
func (db *DB) Exchanges() (es []*Exchange, err error) {
	err = db.Find(&es).Error
	return
}

// ExchangeByName returns an exchange by name
func (db *DB) ExchangeByName(name string) (e Exchange, err error) {
	err = db.Where(&Exchange{Name: name}).First(&e).Error
	return
}
