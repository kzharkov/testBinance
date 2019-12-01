package binance

type AccountResponse struct {
	MakerCommission  int    `json:"makerCommission"`
	TakerCommission  int    `json:"takerCommission"`
	BuyerCommission  int    `json:"buyerCommission"`
	SellerCommission int    `json:"sellerCommission"`
	CanTrade         bool   `json:"canTrade"`
	CanWithdraw      bool   `json:"canWithdraw"`
	CanDeposit       bool   `json:"canDeposit"`
	UpdateTime       int64  `json:"updateTime"`
	AccountType      string `json:"accountType"`
	Balances         []struct {
		Currency string  `json:"currency"`
		Free     float64 `json:"free"`
		Locked   float64 `json:"locked"`
	} `json:"balances"`
}

func (b *Binance) AccountInfo() (*AccountResponse, error) {
	accountResponse := &AccountResponse{}

	err := b.request("/api/v3/account", "GET", nil, accountResponse)
	if err != nil {
		return &AccountResponse{}, err
	}

	return accountResponse, nil
}
