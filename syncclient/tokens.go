package syncclient

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Token struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Address           string `json:"address"`
		Name              string `json:"name"`
		Symbol            string `json:"symbol"`
		ImageUrl          string `json:"image_url"`
		CoingeckoCoinId   string `json:"coingecko_coin_id"`
		Decimals          int    `json:"decimals"`
		TotalSupply       string `json:"total_supply"`
		PriceUsd          string `json:"price_usd"`
		FdvUsd            string `json:"fdv_usd"`
		TotalReserveInUsd string `json:"total_reserve_in_usd"`
		VolumeUsd         struct {
			H24 string `json:"h24"`
		} `json:"volume_usd"`
		MarketCapUsd string `json:"market_cap_usd"`
	} `json:"attributes"`
	Relationships struct {
		TopPools struct {
			Data []struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"top_pools"`
	} `json:"relationships"`
}

type TokenInfo struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Address         string      `json:"address"`
		Name            string      `json:"name"`
		Symbol          string      `json:"symbol"`
		ImageUrl        string      `json:"image_url"`
		CoingeckoCoinId string      `json:"coingecko_coin_id"`
		Websites        []string    `json:"websites"`
		Description     string      `json:"description"`
		GtScore         float64     `json:"gt_score"`
		DiscordUrl      interface{} `json:"discord_url"`
		TelegramHandle  interface{} `json:"telegram_handle"`
		TwitterHandle   interface{} `json:"twitter_handle"`
	} `json:"attributes"`
}

// NetworkTokenPools retrieves the pools for a specific network and token.
// It makes a GET request to the "networks/{network}/tokens/{tokenAddress}/pools/" endpoint of the API.
//
// Parameters:
//   - network: The ID of the network for which to retrieve the pools.
//   - tokenAddress: The address of the token for which to retrieve the pools.
//   - page: The page number for pagination. Each page returns a certain number of pools.
//
// Returns:
//   - A slice of Pool structs, each representing a pool in the network for the given token.
//   - An error if the GET request or the JSON unmarshalling fails.
func (c *Client) NetworkTokenPools(network string, tokenAddress string, page int) (Response[[]Pool], error) {
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	params.Add("include", "base_token,quote_token,dex")

	body, err := c.get(fmt.Sprintf("networks/%s/tokens/%s/pools/", network, tokenAddress), params)
	if err != nil {
		return Response[[]Pool]{}, err
	}
	jsonBody := Response[[]Pool]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return Response[[]Pool]{}, jsonErr
	}
	return jsonBody, nil
}

// NetworkToken retrieves the token for a specific network and address.
// It makes a GET request to the "networks/{network}/tokens/{address}/" endpoint of the API.
//
// Parameters:
//   - network: The ID of the network for which to retrieve the token.
//   - address: The address of the token to be retrieved.
//
// Returns:
//   - A Token struct, representing the token for the given address in the network.
//   - An error if the GET request or the JSON unmarshalling fails.
func (c *Client) NetworkToken(network string, address string) (Response[Token], error) {
	params := url.Values{}
	params.Add("include", "top_pools")

	body, err := c.get(fmt.Sprintf("networks/%s/tokens/%s/", network, address), params)
	if err != nil {
		return Response[Token]{}, err
	}
	jsonBody := Response[Token]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return Response[Token]{}, jsonErr
	}
	return jsonBody, nil
}

// NetworkTokensMultiAddress retrieves the tokens for a specific network and multiple addresses.
// It makes a GET request to the "networks/{network}/tokens/multi/{addresses}" endpoint of the API.
//
// Parameters:
//   - network: The ID of the network for which to retrieve the tokens.
//   - addresses: A slice of addresses for which to retrieve the tokens.
//
// Returns:
//   - A slice of Token structs, each representing a token for the given addresses in the network.
//   - An error if the GET request or the JSON unmarshalling fails.
func (c *Client) NetworkTokensMultiAddress(network string, addresses []string) (Response[[]Token], error) {
	params := url.Values{}
	params.Add("include", "top_pools")

	body, err := c.get(fmt.Sprintf("networks/%s/tokens/multi/%s", network, strings.Join(addresses, ",")), params)
	if err != nil {
		return Response[[]Token]{}, err
	}
	jsonBody := Response[[]Token]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return Response[[]Token]{}, jsonErr
	}
	return jsonBody, nil
}

// NetworkTokenInfo retrieves the token information for a specific network and address.
// It makes a GET request to the "networks/{network}/tokens/{address}/info" endpoint of the API.
//
// Parameters:
//   - network: The ID of the network for which to retrieve the token information.
//   - address: The address of the token for which to retrieve the information.
//
// Returns:
//   - A TokenInfo struct, representing the token information for the given address in the network.
//   - An error if the GET request or the JSON unmarshalling fails.
func (c *Client) NetworkTokenInfo(network string, address string) (Response[TokenInfo], error) {
	body, err := c.get(fmt.Sprintf("networks/%s/tokens/%s/info", network, address), nil)
	if err != nil {
		return Response[TokenInfo]{}, err
	}
	jsonBody := Response[TokenInfo]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return Response[TokenInfo]{}, jsonErr
	}
	return jsonBody, nil
}

// NetworkPoolTokenInfo retrieves the token information for a specific network and pool.
// It makes a GET request to the "networks/{network}/pools/{poolAddress}/info" endpoint of the API.
//
// Parameters:
//   - network: The ID of the network for which to retrieve the token information.
//   - poolAddress: The address of the pool for which to retrieve the token information.
//
// Returns:
//   - A slice of TokenInfo structs, each representing the token information for the given pool in the network.
//   - An error if the GET request or the JSON unmarshalling fails.
func (c *Client) NetworkPoolTokenInfo(network string, poolAddress string) (Response[[]TokenInfo], error) {
	body, err := c.get(fmt.Sprintf("networks/%s/pools/%s/info", network, poolAddress), nil)
	if err != nil {
		return Response[[]TokenInfo]{}, err
	}
	jsonBody := Response[[]TokenInfo]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return Response[[]TokenInfo]{}, jsonErr
	}
	return jsonBody, nil
}

// TokenInfoRecentlyUpdated retrieves the recently updated token information.
// It makes a GET request to the "tokens/info_recently_updated" endpoint of the API.
//
// Returns:
//   - A slice of TokenInfo structs, each representing a recently updated token.
//   - An error if the GET request or the JSON unmarshalling fails.
func (c *Client) TokenInfoRecentlyUpdated() (Response[[]TokenInfo], error) {
	params := url.Values{}
	params.Add("include", "network")
	body, err := c.get("tokens/info_recently_updated", params)
	if err != nil {
		return Response[[]TokenInfo]{}, err
	}
	jsonBody := Response[[]TokenInfo]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return Response[[]TokenInfo]{}, jsonErr
	}
	return jsonBody, nil
}
