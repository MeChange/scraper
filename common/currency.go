// Copyright 2018 MeChange
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"errors"
	"strconv"
)

// A Currency carries a quote, buy rate, and sell rate
type Currency struct {
	Quote             string
	BuyRate, SellRate float64
}

// BuyRateString get buy rate as formatted string
func (curr *Currency) BuyRateString() (string, error) {
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

// SellRateString get sell rate as formatted string
func (curr *Currency) SellRateString() (string, error) {
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
