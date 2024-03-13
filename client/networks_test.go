package client

import (
	"testing"
)

func TestNetworks(t *testing.T) {
	c := NewClient()
	data, _ := c.Networks(1)

	if data[0].Type != "network" {
		t.Error("Expected 'network', got", data[0].Type)
	}
	if data[0].Id != "eth" {
		t.Error("Expected 'eth', got", data[0].Id)
	}
}

func TestNetworkDexes(t *testing.T) {
	c := NewClient()
	data, _ := c.NetworkDexes("eth", 1)

	if data[0].Type != "dex" {
		t.Error("Expected 'dex', got", data[0].Type)
	}
	if data[0].Id != "uniswap_v2" {
		t.Error("Expected 'uniswap_v2', got", data[0].Id)
	}
}

func TestTrendingPools(t *testing.T) {
	c := NewClient()
	data, _ := c.TrendingPools(1)

	if data[0].Type != "pool" {
		t.Error("Expected 'pool', got", data[0].Type)
	}
}

func TestClient_NetworkTrendingPools(t *testing.T) {
	c := NewClient()
	data, _ := c.NetworkTrendingPools("eth", 1)

	if data[0].Type != "pool" {
		t.Error("Expected 'pool', got", data[0].Type)
	}
}
