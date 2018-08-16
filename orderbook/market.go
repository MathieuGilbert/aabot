package orderbook

// Market holds information about tokens per exchange
type Market struct {
	ID         int    `gorm:"primary_key"`
	ExchangeID int    `gorm:"not null"`
	TokenID    int    `gorm:"not null"`
	Balance    string `gorm:"default:'0'"`
	Active     bool   `gorm:"not null"`
}

// Markets returns all the markets
func (db *DB) Markets() (ms []*Market, err error) {
	err = db.Find(&ms).Error
	return
}

// MarketByExTok fetches by exchange id and token id
func (db *DB) MarketByExTok(eid, tid int) (*Market, error) {
	m := &Market{}
	if err := db.Where(&Market{ExchangeID: eid, TokenID: tid}).First(m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
