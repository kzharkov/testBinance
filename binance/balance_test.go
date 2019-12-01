package binance

import (
	"os"
	"testing"
)

func TestBinance_Balance(t *testing.T) {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	type args struct {
		currency string
	}
	tests := []struct {
		name    string
		fields  *Binance
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "Balance BTC",
			fields:  NewBinanceClient(apiKey, secretKey),
			args:    args{},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Binance{
				apiKey:    tt.fields.apiKey,
				secretKey: tt.fields.secretKey,
				client:    tt.fields.client,
			}
			got, err := b.Balance(tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Balance() got = %v, want %v", got, tt.want)
			}
		})
	}
}
