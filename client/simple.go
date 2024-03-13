package client

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TokenPrice struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		TokenPrices map[string]string `json:"token_prices"`
	} `json:"attributes"`
}

func (c *Client) NetworkAddressesTokenPrice(network string, addresses []string) (TokenPrice, error) {
	body, err := c.get(fmt.Sprintf("simple/networks/%s/token_price/%s", network, strings.Join(addresses, ",")), nil)
	if err != nil {
		return TokenPrice{}, err
	}
	jsonBody := response[TokenPrice]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return TokenPrice{}, jsonErr
	}
	return jsonBody.Data, nil
}
