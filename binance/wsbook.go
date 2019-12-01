package binance

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"strings"
)

type bookMsg struct {
	UpdateId int         `json:"lastUpdateId"`
	Bids     [][2]string `json:"bids"`
	Asks     [][2]string `json:"asks"`
}

func (b *Binance) OrderBookStream(ctx context.Context, pair string) (chan OrderBook, error) {
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(BaseWs+strings.ToLower(pair)+"@depth5@100ms", nil)
	if err != nil {
		return nil, err
	}

	ch := make(chan OrderBook, 10)

	go func() {
		select {
		case <-ctx.Done():
			close(ch)
			err = conn.Close()
			if err != nil {
				log.Println(err)
			}
		}
	}()

	go func() {
		for {
			_, binaryMsg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				close(ch)
				return
			}
			msg := &bookMsg{}
			err = json.Unmarshal(binaryMsg, msg)
			if err != nil {
				log.Println(err)
				continue
			}
			orderBook := OrderBook{}
			orderBook.Ask.Amount, err = strconv.ParseFloat(msg.Asks[0][1], 64)
			orderBook.Ask.Price, err = strconv.ParseFloat(msg.Asks[0][0], 64)
			orderBook.Bid.Amount, err = strconv.ParseFloat(msg.Bids[0][1], 64)
			orderBook.Bid.Price, err = strconv.ParseFloat(msg.Bids[0][0], 64)
			if err != nil {
				continue
			}
			ch <- orderBook
		}
	}()

	return ch, nil
}

type OrderBook struct {
	Bid struct {
		Price  float64 `json:"price"`
		Amount float64 `json:"amount"`
	} `json:"bid"`
	Ask struct {
		Price  float64 `json:"price"`
		Amount float64 `json:"amount"`
	} `json:"ask"`
}
