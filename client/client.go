package client

import (
	"io"
	"log"
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

type Response[T Network | Dex | Pool] struct {
	Data  []T   `json:"data"`
	Links Links `json:"Links"`
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
		log.Fatal(err)
	}
	req.Header.Add("accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}
