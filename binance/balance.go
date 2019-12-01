package binance

func (b *Binance) Balance(currency string) (float64, error) {
	accInfo, err := b.AccountInfo()
	if err != nil {
		return 0, err
	}

	for i := range accInfo.Balances {
		if accInfo.Balances[i].Currency == currency {
			return accInfo.Balances[i].Free, nil
		}
	}

	return 0, nil
}
