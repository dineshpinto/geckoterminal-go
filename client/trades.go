package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Trade struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		BlockNumber              int       `json:"block_number"`
		TxHash                   string    `json:"tx_hash"`
		TxFromAddress            string    `json:"tx_from_address"`
		FromTokenAmount          string    `json:"from_token_amount"`
		ToTokenAmount            string    `json:"to_token_amount"`
		PriceFromInCurrencyToken string    `json:"price_from_in_currency_token"`
		PriceToInCurrencyToken   string    `json:"price_to_in_currency_token"`
		PriceFromInUsd           string    `json:"price_from_in_usd"`
		PriceToInUsd             string    `json:"price_to_in_usd"`
		BlockTimestamp           time.Time `json:"block_timestamp"`
		Kind                     string    `json:"kind"`
		VolumeInUsd              string    `json:"volume_in_usd"`
		FromTokenAddress         string    `json:"from_token_address"`
		ToTokenAddress           string    `json:"to_token_address"`
	} `json:"attributes"`
}

func (c *Client) NetworkPoolTrades(network string, poolAddress string, tradeVolumeInUsdGreaterThan int) ([]Trade, error) {
	params := url.Values{}
	params.Add("trade_volume_in_usd_greater_than", strconv.Itoa(tradeVolumeInUsdGreaterThan))
	body, err := c.get(fmt.Sprintf("networks/%s/pools/%s/trades", network, poolAddress), params)
	if err != nil {
		return nil, err
	}
	jsonBody := response[[]Trade]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return jsonBody.Data, nil
}
