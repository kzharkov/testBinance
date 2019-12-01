package binance

import (
	"context"
	"os"
	"testing"
)

func TestBinance_OrderBookStream(t *testing.T) {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	type fields struct {
		apiKey    string
		secretKey string
	}
	type args struct {
		ctx  context.Context
		pair string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    chan OrderBook
		wantErr bool
	}{
		{
			name: "BTC_USDT",
			fields: fields{
				apiKey:    apiKey,
				secretKey: secretKey,
			},
			args: args{
				ctx:  context.Background(),
				pair: "btcusd",
			},
			want:    make(chan OrderBook),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBinanceClient(tt.fields.apiKey, tt.fields.secretKey)
			_, err := b.OrderBookStream(tt.args.ctx, tt.args.pair)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderBookStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
