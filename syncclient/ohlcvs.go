package syncclient

import (
	"encoding/json"
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"time"
)

type OHLCVS struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		OHLCVList [][]float64 `json:"ohlcv_list"`
	} `json:"attributes"`
}

type NetworkPoolOHLCVArgs struct {
	Network     string
	PoolAddress string
	Timeframe   string
	// Optional parameters
	Aggregate       int
	BeforeTimestamp int64
	Limit           int
	Currency        string
	Token           string
}

// NetworkPoolOHLCV retrieves the OHLCV (Open, High, Low, Close, Volume) data for a specific pool in a network.
// It makes a GET request to the "networks/{network}/pools/{poolAddress}/ohlcv/{timeframe}" endpoint of the API.
//
// Parameters:
//   - args: A struct of NetworkPoolOHLCVArgs which includes:
//   - Network: The ID of the network for which to retrieve the OHLCV data.
//   - PoolAddress: The address of the pool for which to retrieve the OHLCV data.
//   - Timeframe: The timeframe for the OHLCV data ("minute", "hour", "day").
//   - Aggregate: The aggregate level for the OHLCV data 1, 5, 15 for minute, 1, 4, 12 for hour, 1 for day (default 1).
//   - BeforeTimestamp: The Unix timestamp before which to retrieve the OHLCV data (default now).
//   - Limit: The number of OHLCVs to return (default 100).
//   - Currency: The currency for the OHLCV data ("usd", "token") (default "usd").
//   - Token: The token for the OHLCV data ("base", "quote") (default "base").
//
// Returns:
//   - An OHLCVS struct, representing the OHLCV data for the given pool in the network.
//   - An error if the GET request or the JSON unmarshalling fails.
func (c *Client) NetworkPoolOHLCV(args NetworkPoolOHLCVArgs) (OHLCVS, error) {
	if args.Network == "" {
		return OHLCVS{}, fmt.Errorf("network is required")
	}
	if args.PoolAddress == "" {
		return OHLCVS{}, fmt.Errorf("pool address is required")
	}
	if args.Timeframe == "" {
		return OHLCVS{}, fmt.Errorf("timeframe is required")
	}

	if args.Aggregate == 0 {
		args.Aggregate = 1
	} else {
		if args.Timeframe == "day" {
			validDayAggregates := []int{1}
			if !slices.Contains(validDayAggregates, args.Aggregate) {
				return OHLCVS{}, fmt.Errorf("invalid aggregate for day timeframe")
			}
		} else if args.Timeframe == "minute" {
			validMinuteAggregates := []int{1, 5, 15}
			if !slices.Contains(validMinuteAggregates, args.Aggregate) {
				return OHLCVS{}, fmt.Errorf("invalid aggregate for minute timeframe")
			}
		} else if args.Timeframe == "hour" {
			validHourAggregates := []int{1, 4, 12}
			if !slices.Contains(validHourAggregates, args.Aggregate) {
				return OHLCVS{}, fmt.Errorf("invalid aggregate for hour timeframe")
			}
		} else {
			return OHLCVS{}, fmt.Errorf("invalid timeframe")
		}
	}
	if args.BeforeTimestamp == 0 {
		args.BeforeTimestamp = time.Now().Unix()
	}
	if args.Limit == 0 {
		args.Limit = 100
	}
	if args.Currency == "" {
		args.Currency = "usd"

	}
	if args.Token == "" {
		args.Token = "base"
	}

	params := url.Values{}
	params.Add("aggregate", strconv.Itoa(args.Aggregate))
	params.Add("before_timestamp", strconv.FormatInt(args.BeforeTimestamp, 10))
	params.Add("limit", strconv.Itoa(args.Limit))
	params.Add("currency", args.Currency)
	params.Add("token", args.Token)

	body, err := c.get(fmt.Sprintf("networks/%s/pools/%s/ohlcv/%s", args.Network, args.PoolAddress, args.Timeframe), params)
	if err != nil {
		return OHLCVS{}, err

	}

	jsonBody := responseOHLCVS{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return OHLCVS{}, jsonErr
	}
	return jsonBody.Data, nil
}
