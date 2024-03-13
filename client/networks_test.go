package client

import (
	"testing"
)

func TestClient_Networks(t *testing.T) {
	gt := GeckoTerminalClient()
	data, _ := gt.Networks(1)

	if data[0].Type != "network" {
		t.Error("Expected 'network', got", data[0].Type)
	}
	if data[0].Id != "eth" {
		t.Error("Expected 'eth', got", data[0].Id)
	}
}

func TestClient_NetworkDexes(t *testing.T) {
	gt := GeckoTerminalClient()
	data, _ := gt.NetworkDexes("eth", 1)

	if data[0].Type != "dex" {
		t.Error("Expected 'dex', got", data[0].Type)
	}
	if data[0].Id != "uniswap_v2" {
		t.Error("Expected 'uniswap_v2', got", data[0].Id)
	}
}

func TestClient_TrendingPools(t *testing.T) {
	gt := GeckoTerminalClient()
	data, _ := gt.TrendingPools(1)

	if data[0].Type != "pool" {
		t.Error("Expected 'pool', got", data[0].Type)
	}
}

func TestClient_NetworkTrendingPools(t *testing.T) {
	gt := GeckoTerminalClient()
	data, _ := gt.NetworkTrendingPools("eth", 1)

	if data[0].Type != "pool" {
		t.Error("Expected 'pool', got", data[0].Type)
	}
}

func TestClient_NetworkPoolAddress(t *testing.T) {
	gt := GeckoTerminalClient()
	data, _ := gt.NetworkPoolAddress("eth", "0x60594a405d53811d3bc4766596efd80fd545a270")

	if data.Type != "pool" {
		t.Error("Expected 'pool', got", data.Type)
	}
	if data.Id != "eth_0x60594a405d53811d3bc4766596efd80fd545a270" {
		t.Error("Expected 'eth_0x60594a405d53811d3bc4766596efd80fd545a270', got", data.Id)
	}
}

func TestClient_NetworkPoolMultiAddress(t *testing.T) {
	gt := GeckoTerminalClient()
	addresses := []string{"0x60594a405d53811d3bc4766596efd80fd545a270", "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"}
	data, _ := gt.NetworkPoolMultiAddress("eth", addresses)

	if data[0].Type != "pool" {
		t.Error("Expected 'pool', got", data[0].Type)
	}
	if data[0].Id != "eth_0x60594a405d53811d3bc4766596efd80fd545a270" {
		t.Error("Expected 'eth_0x60594a405d53811d3bc4766596efd80fd545a270', got", data[0].Id)
	}
	if data[1].Type != "pool" {
		t.Error("Expected 'pool', got", data[1].Type)
	}
	if data[1].Id != "eth_0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640" {
		t.Error("Expected 'eth_0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640', got", data[1].Id)
	}
}

func TestClient_NetworkPools(t *testing.T) {
	gt := GeckoTerminalClient()
	data, _ := gt.NetworkPools("eth", 1)

	if data[0].Type != "pool" {
		t.Error("Expected 'pool', got", data[0].Type)
	}
	if data[0].Id != "eth_0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640" {
		t.Error("Expected 'eth_0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640', got", data[0].Id)
	}
}
