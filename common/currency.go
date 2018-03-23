// Copyright 2018 Pongpat Weesommai. All rights reserved.

package common

import (
	"errors"
	"strconv"
)

// A Currency carries a quote and a buy and sell rate
type Currency struct {
	Quote string
	BuyRate, SellRate float64
}

func (curr *Currency) GetBuyRateString() (string, error) {
	if curr.BuyRate <= 0 {
		return "", errors.New("Buy rate is less than or equal to zero") 
	}
	// For currency that rate is >= 1.0 has 2 decimal points e.g. USD
	if curr.BuyRate >= 1.0 {
		return strconv.FormatFloat(curr.BuyRate, 'f', 2, 64), nil
	}
	// For currency that rate is < 1.0 and > 0 has 4 decimal points e.g. JPY
	return strconv.FormatFloat(curr.BuyRate, 'f', 4, 64), nil
}

func (curr *Currency) GetSellRateString() (string, error) {
	if curr.SellRate <= 0 {
		return "", errors.New("Sell rate is less than or equal to zero")
	}
	// For currency that rate is >= 1.0 has 2 decimal points e.g. USD
	if curr.SellRate >= 1.0 {
		return strconv.FormatFloat(curr.SellRate, 'f', 2, 64), nil
	}
	// For currency that rate is < 1.0 and > 0 has 4 decimal points e.g. JPY
	return strconv.FormatFloat(curr.SellRate, 'f', 4, 64), nil
}