package syncclient

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Network struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name                     string `json:"name"`
		CoingeckoAssetPlatformId string `json:"coingecko_asset_platform_id"`
	} `json:"attributes"`
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
func (c *Client) Networks(page int) (Response[[]Network], error) {
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	body, err := c.get("networks/", params)
	if err != nil {
		return Response[[]Network]{}, err
	}
	jsonBody := Response[[]Network]{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return Response[[]Network]{}, jsonErr
	}
	return jsonBody, nil
}
