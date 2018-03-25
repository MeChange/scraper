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