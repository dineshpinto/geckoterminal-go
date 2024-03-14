package syncclient

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

// NetworkAddressesTokenPrice retrieves the token price for a specific network and addresses.
// It makes a GET request to the "simple/networks/{network}/token_price/{addresses}" endpoint of the API.
//
// Parameters:
//   - network: The ID of the network for which to retrieve the token price.
//   - addresses: A slice of addresses for which to retrieve the token price.
//
// Returns:
//   - A TokenPrice struct, representing the token price for the given addresses in the network.
//   - An error if the GET request or the JSON unmarshalling fails.
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
