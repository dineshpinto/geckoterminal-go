package main

import (
	"geckoterminal-go/client"
)

func main() {
	c := client.GeckoTerminalClient()
	c.NetworkTrendingPools("eth", 1)
}
