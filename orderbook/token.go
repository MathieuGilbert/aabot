package orderbook

// Token holds the name-address pair
type Token struct {
	ID      int    `gorm:"primary_key" json:"id"`
	Name    string `gorm:"not null;index" json:"name"`
	Address string `gorm:"not null" json:"address"`
}

// Tokens returns all the tokens
func (db *DB) Tokens() (ts []*Token, err error) {
	err = db.Find(&ts).Order("id").Error
	return
}

// Token returns the token with the id
func (db *DB) Token(id int) (*Token, error) {
	t := &Token{}
	err := db.First(t, id).Error
	return t, err
}

// TokenByName returns a token by name
func (db *DB) TokenByName(name string) (t Token, err error) {
	err = db.Where(&Token{Name: name}).First(&t).Error
	return
}
