package orderbook

import (
	"bytes"
	"encoding/json"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	goed "github.com/mathieugilbert/go-etherdelta"
)

// Exchange holds exchange-specific details
type Exchange struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `gorm:"not null;index" json:"name"`
	Address      string    `gorm:"not null" json:"address"`
	WebsocketURI string    `gorm:"not null" json:"websocket_uri"`
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
		log.Printf("Error getting markets for exchange %v: %v\n", e.Name, err)
		return err
	}

	// going to wait until all markets finish
	var wg sync.WaitGroup
	wg.Add(len(ms))

	for _, m := range ms {
		go m.LoadOrderbook(db, s, &wg)
	}

	wg.Wait()
	return nil
}

// SyncOrderbook connects to exchange websocket and updates orderbook as messages arrive
func (e *Exchange) SyncOrderbook(db *DB, s *ExchangeServices, interrupt chan os.Signal) error {
	log.Printf("Connecting to %v websocket...", e.Name)

	url := "wss://api.forkdelta.com/socket.io/?EIO=3&transport=websocket"
	c, resp, err := websocket.DefaultDialer.Dial(url, nil)
	if err == websocket.ErrBadHandshake {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		log.Printf("%+v\n", buf.String())
		log.Printf("handshake failed with status %d", resp.StatusCode)
	}
	if err != nil {
		log.Panicf("Dial(%v) error: %v\n", url, err)
	}
	defer c.Close()

	log.Printf("Connected to %v websocket.", e.Name)

	done := make(chan struct{})

	var respRxp = regexp.MustCompile(`(\d+)(.*)`)

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if err.Error() == "websocket: close 1000 (normal)" {
					log.Println("Websocket closed")
					break
				} else {
					log.Println("ReadMessage error:", err)
					continue
				}
			}

			match := respRxp.FindStringSubmatch(string(message))
			if len(match) != 3 {
				log.Printf("Regex didn't find matches in: %v", message)
				continue
			}

			switch match[1] {
			case "0":
				log.Printf("Open: %v\n", match[2])
			case "40":
				log.Println("Connected to websocket")
			case "42":
				var orderRxp = regexp.MustCompile(`\[.(.*?).,(.*)]`)
				matchOrder := orderRxp.FindStringSubmatch(match[2])
				if len(matchOrder) != 3 {
					log.Printf("42, not orders: %v\n", match[2])
					break
				}

				switch matchOrder[1] {
				case "orders":
					resp := &goed.OrderBook{}

					if err = json.Unmarshal([]byte(matchOrder[2]), resp); err != nil {
						log.Printf("Can't unmarshal orders: %v\n%v\n", matchOrder[2], err)
						break
					}

					count, err := processNewOrders(db, e, resp)
					if err != nil {
						log.Printf("Unable to process new orders: %v\n", err)
						break
					}
					if count > 0 {
						log.Printf("Inserted %v new orders\n", count)
					}
				case "trades":
					type Trade struct {
						TxHash     string `json:"txHash"`
						Date       string `json:"date"`
						Price      string `json:"price"`
						Side       string `json:"side"`
						Amount     string `json:"amount"`
						AmountBase string `json:"amountBase"`
						Buyer      string `json:"buyer"`
						Seller     string `json:"seller"`
						TokenAddr  string `json:"tokenAddr"`
					}
					var trades []*Trade

					if err = json.Unmarshal([]byte(matchOrder[2]), &trades); err != nil {
						log.Panicf("Can't unmarshal trades: %v\n%v\n", matchOrder[2], err)
						break
					}

					for _, trade := range trades {
						t, err := db.TokenByAddress(trade.TokenAddr)
						if err != nil {
							log.Printf("Don't care about %v token\n", trade.TokenAddr)
							continue
						}
						// re-fetch orderbook for token
						m, err := db.MarketByExTok(e.ID, t.ID)
						if err != nil {
							log.Printf("Error fetching market for exchange %v, token %v\n", e.ID, t.ID)
							continue
						}

						var wg sync.WaitGroup
						wg.Add(1)
						go func(m *Market) {
							m.LoadOrderbook(db, s, &wg)
						}(m)
						wg.Wait()

						log.Printf("New trades, refreshed orderbook for token %v on exchange %v\n", t.Name, e.Name)
					}

				}
			default:
				log.Printf("other message: %v\n", match[0])
			}
		}

	}()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return nil
		case <-ticker.C:
			err := c.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				log.Println("WriteMessage error:", err)
				return err
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error writing close message:", err)
				return err
			}
			select {
			case <-done:
				log.Println("DONE")
			case <-time.After(time.Second):
				log.Println("A SECOND")
			}
			return nil
		}
	}
}

func processNewOrders(db *DB, e *Exchange, ob *goed.OrderBook) (int, error) {
	var volumeThreshold = 0.001 * math.Pow10(18)

	var os []*Order

	for _, order := range ob.Buys {
		vol, err := strconv.ParseFloat(order.AvailableVolumeBase, 32)
		if err != nil {
			log.Printf("ParseFloat(%v) error: %v\n", order.AvailableVolumeBase, err)
			continue
		}
		if vol < volumeThreshold {
			continue
		}

		t, err := db.TokenByAddress(order.TokenGet)
		if err != nil {
			// if not in db, we don't care about it
			continue
		}

		m, err := db.MarketByExTok(e.ID, t.ID)
		if err != nil {
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

	for _, order := range ob.Sells {
		vol, err := strconv.ParseFloat(order.AvailableVolumeBase, 32)
		if err != nil {
			log.Printf("ParseFloat(%v) error: %v\n", order.AvailableVolumeBase, err)
			continue
		}
		if vol < volumeThreshold {
			continue
		}

		t, err := db.TokenByAddress(order.TokenGive)
		if err != nil {
			// if not in db, we don't care about it
			continue
		}

		m, err := db.MarketByExTok(e.ID, t.ID)
		if err != nil {
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

	if len(os) == 0 {
		return 0, nil
	}

	if err := db.BulkInsertOrders(os); err != nil {
		return 0, err
	}
	return len(os), nil
}
