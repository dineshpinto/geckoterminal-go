package main

import (
	"geckoterminal-go/client"
)

func main() {
	c := client.NewClient()
	c.NetworkTrendingPools("eth", 1)
}
