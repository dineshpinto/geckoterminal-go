package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Network struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name                     string `json:"name"`
		CoingeckoAssetPlatformId string `json:"coingecko_asset_platform_id"`
	} `json:"attributes"`
}

type Dex struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name string `json:"name"`
	} `json:"attributes"`
}

type Pool struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		BaseTokenPriceUsd             string      `json:"base_token_price_usd"`
		BaseTokenPriceNativeCurrency  string      `json:"base_token_price_native_currency"`
		QuoteTokenPriceUsd            string      `json:"quote_token_price_usd"`
		QuoteTokenPriceNativeCurrency string      `json:"quote_token_price_native_currency"`
		BaseTokenPriceQuoteToken      string      `json:"base_token_price_quote_token"`
		QuoteTokenPriceBaseToken      string      `json:"quote_token_price_base_token"`
		Address                       string      `json:"address"`
		Name                          string      `json:"name"`
		PoolCreatedAt                 time.Time   `json:"pool_created_at"`
		FdvUsd                        string      `json:"fdv_usd"`
		MarketCapUsd                  interface{} `json:"market_cap_usd"`
		PriceChangePercentage         struct {
			M5  string `json:"m5"`
			H1  string `json:"h1"`
			H6  string `json:"h6"`
			H24 string `json:"h24"`
		} `json:"price_change_percentage"`
		Transactions struct {
			M5 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"m5"`
			M15 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"m15"`
			M30 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"m30"`
			H1 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"h1"`
			H24 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"h24"`
		} `json:"transactions"`
		VolumeUsd struct {
			M5  string `json:"m5"`
			H1  string `json:"h1"`
			H6  string `json:"h6"`
			H24 string `json:"h24"`
		} `json:"volume_usd"`
		ReserveInUsd string `json:"reserve_in_usd"`
	} `json:"attributes"`
	Relationships struct {
		BaseToken struct {
			Data struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"base_token"`
		QuoteToken struct {
			Data struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"quote_token"`
		Network struct {
			Data struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"Network"`
		Dex struct {
			Data struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"Dex"`
	} `json:"relationships"`
}

// Networks retrieves the list of networks.
// It makes a GET request to the "networks/" endpoint of the API.
//
// Parameters:
//   - page: The page number for pagination. Each page returns a certain number of networks.
//
// Returns:
//   - A slice of Network structs, each representing a network.
//   - An error if the GET request or the JSON unmarshalling fails.
func (c *Client) Networks(page int) ([]Network, error) {
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	body, err := c.get("networks/", params)
	if err != nil {
		return nil, err
	}
	jsonBody := response[[]Network]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return jsonBody.Data, nil
}

// NetworkDexes retrieves the Dexes for a specific network.
// It makes a GET request to the "networks/{network}/dexes/" endpoint of the API.
//
// Parameters:
//   - network: The ID of the network for which to retrieve the Dexes.
//   - page: The page number for pagination. Each page returns a certain number of Dexes.
//
// Returns:
//   - A slice of Dex structs, each representing a Dex in the network.
//   - An error if the GET request or the JSON unmarshalling fails.
func (c *Client) NetworkDexes(network string, page int) ([]Dex, error) {
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	body, err := c.get(fmt.Sprintf("networks/%s/dexes/", network), params)
	if err != nil {
		return nil, err
	}
	jsonBody := response[[]Dex]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return jsonBody.Data, nil
}

// TrendingPools retrieves the trending pools across all networks.
// It makes a GET request to the "networks/trending_pools/" endpoint of the API.
//
// Parameters:
//   - page: The page number for pagination. Each page returns a certain number of trending pools.
//
// Returns:
//   - A slice of Pool structs, each representing a trending Pool across all networks.
//   - An error if the GET request or the JSON unmarshalling fails.
func (c *Client) TrendingPools(page int) ([]Pool, error) {
	params := url.Values{}
	params.Add("include", "base_token,quote_token,dex,network")
	params.Add("page", strconv.Itoa(page))
	body, err := c.get("networks/trending_pools/", params)
	if err != nil {
		return nil, err
	}
	jsonBody := response[[]Pool]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return jsonBody.Data, nil
}

// NetworkTrendingPools retrieves the trending pools for a specific Network.
// It makes a GET request to the "networks/{Network}/trending_pools/" endpoint of the API.
//
// Parameters:
//   - Network: The ID of the Network for which to retrieve the trending pools.
//   - page: The page number for pagination. Each page returns a certain number of trending pools.
//
// Returns:
//   - A slice of Pool structs, each representing a trending Pool in the Network.
//   - An error if the GET request or the JSON unmarshalling fails.
func (c *Client) NetworkTrendingPools(network string, page int) ([]Pool, error) {
	params := url.Values{}
	params.Add("include", "base_token,quote_token,dex")
	params.Add("page", strconv.Itoa(page))
	body, err := c.get(fmt.Sprintf("networks/%s/trending_pools/", network), params)
	if err != nil {
		return nil, err
	}
	jsonBody := response[[]Pool]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return jsonBody.Data, nil
}

func (c *Client) NetworkPoolAddress(network string, address string) (Pool, error) {
	params := url.Values{}
	params.Add("include", "base_token,quote_token,dex")
	body, err := c.get(fmt.Sprintf("networks/%s/pools/%s", network, address), params)
	if err != nil {
		return Pool{}, err
	}
	jsonBody := response[Pool]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return Pool{}, jsonErr
	}
	return jsonBody.Data, nil
}

func (c *Client) NetworkPoolMultiAddress(network string, address []string) ([]Pool, error) {
	params := url.Values{}
	params.Add("include", "base_token,quote_token,dex")
	body, err := c.get(fmt.Sprintf("networks/%s/pools/multi/%s", network, strings.Join(address, ",")), params)
	if err != nil {
		return nil, err
	}
	jsonBody := response[[]Pool]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return jsonBody.Data, nil
}

func (c *Client) NetworkPools(network string, page int) ([]Pool, error) {
	params := url.Values{}
	params.Add("include", "base_token,quote_token,dex")
	params.Add("page", strconv.Itoa(page))
	body, err := c.get(fmt.Sprintf("networks/%s/pools/", network), params)
	if err != nil {
		return nil, err
	}
	jsonBody := response[[]Pool]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return jsonBody.Data, nil
}
