package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Dex struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name string `json:"name"`
	} `json:"attributes"`
}

// NetworkDexes retrieves the Dexes for a specific network.
// It makes a GET request to the "networks/{network}/dexes.go/" endpoint of the API.
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
	body, err := c.get(fmt.Sprintf("networks/%s/dexes.go/", network), params)
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
