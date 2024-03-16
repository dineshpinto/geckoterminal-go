[![Go](https://github.com/dineshpinto/geckoterminal-go/actions/workflows/go.yml/badge.svg)](https://github.com/dineshpinto/geckoterminal-go/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/dineshpinto/geckoterminal-go/graph/badge.svg?token=PPR56W2NSI)](https://codecov.io/gh/dineshpinto/geckoterminal-go)

# GeckoTerminal Go

## RESTful Go Client for GeckoTerminal API

Wrapper around the [GeckoTerminal](https://www.geckoterminal.com) DeFi and DeX
aggregator operating across 90+ chains and 500+ dexes.

Features:

- Get the market data (price, volume, historical chart) of any token
- Find all the pools that are trading a specific token
- Plot a candlestick chart using OHLCV when given a pool address

The API is currently in beta and is subject to change, please report any issues you
find.

## Installation

```bash
go get github.com/dineshpinto/geckoterminal-go
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/dineshpinto/geckoterminal-go/syncclient"
)

func main() {
	// Create an instance of the synchronous client
	gt := syncclient.NewClient()

	// Get list of networks
	networks, err := gt.Networks(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(networks)

	// Get list of trending pools on a network
	pools, err := gt.NetworkTrendingPools("solana", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pools)

	// and many more...
}
```

## Examples

### Get list of supported networks

```go
networks, _ := gt.Networks(1)

for _, network := range networks.Data {
    fmt.println(network.Id, network.Attributes.CoingeckoAssetPlatformId)
}

```

Output:
```text
eth ethereum
bsc binance-smart-chain
polygon_pos polygon-pos
avax avalanche
...
```

### Get list of trending pools on a network

```go
trending, _ := gt.NetworkTrendingPools("solana", 1)

for _, pool := range trending.Data {
    fmt.println(pool.Attributes.Name, pool.Attributes.BaseTokenPriceUsd, pool.Attributes.PriceChangePercentage.H24, pool.Attributes.VolumeUsd.H24)
}
```

Output:
```text
BOME / SOL 0.0118312250522804 200.03 398924901.536168
MONKEY / SOL 0.00000000233851081107242 422.12 328827.649694968
PORTNOY / SOL 0.00113291211384906 -58.77 1751325.64299133
PENG / SOL 0.850157764212865 -54.64 14786553.562191
NINJA / SOL 0.0128840288743118 -2.39 2749158.88490154
boden / SOL 0.130505062890272 -17.74 4908971.2246087
$WIF / SOL 2.5686889316078 -20.82 30711940.7876968
...
```

## Disclaimer

This project is for educational purposes only. You should not construe any such
information or other material as legal, tax, investment, financial, or other advice.
Nothing contained here constitutes a solicitation, recommendation, endorsement, or
offer by me or any third party service provider to buy or sell any securities or other
financial instruments in this or in any other jurisdiction in which such solicitation or
offer would be unlawful under the securities laws of such jurisdiction.

Under no circumstances will I be held responsible or liable in any way for any claims,
damages, losses, expenses, costs, or liabilities whatsoever, including, without
limitation, any direct or indirect damages for loss of profits.