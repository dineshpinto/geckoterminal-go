[![Go](https://github.com/dineshpinto/geckoterminal-go/actions/workflows/go.yml/badge.svg)](https://github.com/dineshpinto/geckoterminal-go/actions/workflows/go.yml)

# GeckoTerminal Go

## RESTful Go Client for GeckoTerminal API

Wrapper around the [GeckoTerminal](https://www.geckoterminal.com) DeFi and DeX
aggregator operating across 90+ chains and 500+ dexes.
Features both synchronous and asynchronous APIs.

Features:

- Get the market data (price, volume, historical chart) of any token
- Find all the pools that are trading a specific token
- Plot a candlestick chart using OHLCV when given a pool address

The API is currently in beta and is subject to change, please report any issues you
find.

## Installation

```bash
go install github.com/dineshpinto/geckoterminal-go
```

## Examples

```go
package main

import (
	"fmt"
	"log"

	"github.com/dineshpinto/geckoterminal-go/syncclient"
)

func main() {
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
}
```