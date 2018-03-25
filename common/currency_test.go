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
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetBuyRateString(t *testing.T) {
	curr := Currency{Quote: "USD", BuyRate: -1, SellRate: -1}
	buyRate, err := curr.GetBuyRateString()
	assert.Equal(t, "", buyRate, "they should be equal")
	assert.NotNil(t, err)

	curr = Currency{Quote: "EUR", BuyRate: 37.83375, SellRate: 39.18625}
	buyRate, err = curr.GetBuyRateString()
	assert.Equal(t, "37.83", buyRate, "they should be equal")
	assert.Nil(t, err)

	curr = Currency{Quote: "JPY", BuyRate: 0.2934875, SellRate: 0.3020250}
	buyRate, err = curr.GetBuyRateString()
	assert.Equal(t, "0.2935", buyRate, "they should be equal")
	assert.Nil(t, err)
}

func TestGetSellRateString(t *testing.T) {
	curr := Currency{Quote: "USD", BuyRate: -1, SellRate: -1}
	sellRate, err := curr.GetSellRateString()
	assert.Equal(t, "", sellRate, "they should be equal")
	assert.NotNil(t, err)

	curr = Currency{Quote: "EUR", BuyRate: 37.83375, SellRate: 39.18625}
	sellRate, err = curr.GetSellRateString()
	assert.Equal(t, "39.19", sellRate, "they should be equal")
	assert.Nil(t, err)

	curr = Currency{Quote: "JPY", BuyRate: 0.2934875, SellRate: 0.3020250}
	sellRate, err = curr.GetSellRateString()
	assert.Equal(t, "0.3020", sellRate, "they should be equal")
	assert.Nil(t, err)
}