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
    fmt.Printf("%s: (%s, %s)\n",
        network.Attributes.Name, network.Id, network.Attributes.CoingeckoAssetPlatformId,
    )
}

```

Output:
```text
Ethereum: (eth, ethereum)
Solana: (solana, solana)
Arbitrum: (arbitrum, arbitrum-one)
Polygon POS: (polygon_pos, polygon-pos)
Avalanche: (avax, avalanche)
Mantle: (mantle, mantle)
...
```

### Get list of trending pools on a network

```go
trending, _ := gt.NetworkTrendingPools("solana", 1)

for _, pool := range trending.Data {
    fmt.Printf("%s: $%s (24H: %s%%), Vol: $%s, %s (%s)\n",
        pool.Attributes.Name, pool.Attributes.BaseTokenPriceUsd,
        pool.Attributes.PriceChangePercentage.H24, pool.Attributes.VolumeUsd.H24,
        pool.Attributes.Address, pool.Relationships.Dex.Data.Id,
    )
}
```

Output:
```text
BOME / SOL: $0.0126440024306716 (24H: 211.42%), Vol: $395895571.566249, DSUvc5qf5LJHHV5e2tD184ixotSnCnwj7i4jJa4Xsrmt (raydium)
MONKEY / SOL: $0.00000000229179523969613 (24H: 394.16%), Vol: $327259.682046244, Dqb7bL7MZkuDgHrZZphRMRViJnepHxf9odx3RRwmifur (raydium)
PORTNOY / SOL: $0.00106720074955826 (24H: -65.69%), Vol: $1716765.08119867, 77JrcxAzPUEvn9o1YXmFm9zQid8etT4SCWVxVqE8VTTG (raydium)
PENG / SOL: $0.892427610912417 (24H: -51.96%), Vol: $14666102.7543845, AxBDdiMK9hRPLMPM7k6nCPC1gRARgXQHNejfP2LvNGr6 (raydium)
NINJA / SOL: $0.0131834602018283 (24H: 7.1%), Vol: $2725307.42747332, B8sv1wiDf9VydktLmBDHUn4W5UFkFAcbgu7igK6Fn2sW (raydium)
boden / SOL: $0.124381368872888 (24H: -22.81%), Vol: $4945908.16723495, 6UYbX1x8YUcFj8YstPYiZByG7uQzAq2s46ZWphUMkjg5 (raydium)
$WIF / SOL: $2.52566860102974 (24H: -21.22%), Vol: $30680222.6561265, EP2ib6dYdEeqD8MfE2ezHCxX3kP3K2eLKkirfPm5eyMx (raydium)
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