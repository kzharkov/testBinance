package main

import (
	"context"
	"darwin/test/binance"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	var client Exchange
	client = binance.NewBinanceClient(apiKey, secretKey)

	balanceBTC, err := client.Balance("BTC")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("BTC balance = ", balanceBTC)

	ctx, cancel := context.WithCancel(context.Background())

	orderBookStream, err := client.OrderBookStream(ctx, "btcusdt")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("OrderBook btcusdt")
	for i := 0; i < 5; i++ {
		orderBookBTCUSDT := <-orderBookStream
		output, err := json.Marshal(orderBookBTCUSDT)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(string(output))
	}
	cancel()
}

type Exchange interface {
	Balance(currency string) (float64, error)
	OrderBookStream(ctx context.Context, pair string) (chan binance.OrderBook, error)
}
