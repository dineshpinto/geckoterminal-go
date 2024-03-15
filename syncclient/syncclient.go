package syncclient

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	baseUrl    string
	httpClient *http.Client
}

type Links struct {
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

type Meta struct {
	Base struct {
		Address         string `json:"address"`
		Name            string `json:"name"`
		Symbol          string `json:"symbol"`
		CoingeckoCoinId string `json:"coingecko_coin_id"`
	} `json:"base"`
	Quote struct {
		Address         string `json:"address"`
		Name            string `json:"name"`
		Symbol          string `json:"symbol"`
		CoingeckoCoinId string `json:"coingecko_coin_id"`
	} `json:"quote"`
}

type DataTypes interface {
	[]Network | []Dex | []Pool | []Trade | []Token | Pool | TokenPrice | Token | []TokenInfo | TokenInfo | OHLCVS
}

type Response[T DataTypes] struct {
	Data  T     `json:"data"`
	Links Links `json:"links"`
	Meta  Meta  `json:"meta"`
}

func NewClient() *Client {
	return &Client{
		baseUrl:    "https://api.geckoterminal.com/api/v2/",
		httpClient: http.DefaultClient,
	}
}

func (c *Client) get(endpoint string, params url.Values) ([]byte, error) {
	req, err := http.NewRequest("GET", c.baseUrl+endpoint, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
