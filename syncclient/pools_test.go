package syncclient

import "testing"

func TestClient_TrendingPools(t *testing.T) {
	gt := NewClient()
	resp, _ := gt.TrendingPools(1)

	if resp.Data[0].Type != "pool" {
		t.Error("Expected 'pool', got", resp.Data[0].Type)
	}
}

func TestClient_NetworkTrendingPools(t *testing.T) {
	gt := NewClient()
	resp, _ := gt.NetworkTrendingPools("eth", 1)

	if resp.Data[0].Type != "pool" {
		t.Error("Expected 'pool', got", resp.Data[0].Type)
	}
}

func TestClient_NetworkPoolAddress(t *testing.T) {
	gt := NewClient()
	resp, _ := gt.NetworkPoolAddress("eth", "0x60594a405d53811d3bc4766596efd80fd545a270")

	if resp.Data.Type != "pool" {
		t.Error("Expected 'pool', got", resp.Data.Type)
	}
	if resp.Data.Id != "eth_0x60594a405d53811d3bc4766596efd80fd545a270" {
		t.Error("Expected 'eth_0x60594a405d53811d3bc4766596efd80fd545a270', got", resp.Data.Id)
	}
}

func TestClient_NetworkPoolMultiAddress(t *testing.T) {
	gt := NewClient()
	addresses := []string{"0x60594a405d53811d3bc4766596efd80fd545a270", "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"}
	resp, _ := gt.NetworkPoolMultiAddress("eth", addresses)

	if resp.Data[0].Type != "pool" {
		t.Error("Expected 'pool', got", resp.Data[0].Type)
	}
	if resp.Data[0].Id != "eth_0x60594a405d53811d3bc4766596efd80fd545a270" {
		t.Error("Expected 'eth_0x60594a405d53811d3bc4766596efd80fd545a270', got", resp.Data[0].Id)
	}
	if resp.Data[1].Type != "pool" {
		t.Error("Expected 'pool', got", resp.Data[1].Type)
	}
	if resp.Data[1].Id != "eth_0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640" {
		t.Error("Expected 'eth_0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640', got", resp.Data[1].Id)
	}
}

func TestClient_NetworkPools(t *testing.T) {
	gt := NewClient()
	resp, _ := gt.NetworkPools("eth", 1)

	if resp.Data[0].Type != "pool" {
		t.Error("Expected 'pool', got", resp.Data[0].Type)
	}
}

func TestClient_NetworkDexPools(t *testing.T) {
	gt := NewClient()
	resp, _ := gt.NetworkDexPools("eth", "sushiswap", 1)

	if resp.Data[0].Type != "pool" {
		t.Error("Expected 'pool', got", resp.Data[0].Type)
	}
	if resp.Data[0].Id != "eth_0x397ff1542f962076d0bfe58ea045ffa2d347aca0" {
		t.Error("Expected 'eth_0x397ff1542f962076d0bfe58ea045ffa2d347aca0', got", resp.Data[0].Id)
	}
}

func TestClient_NetworkNewPools(t *testing.T) {
	gt := NewClient()
	resp, _ := gt.NetworkNewPools("eth", 1)

	if resp.Data[0].Type != "pool" {
		t.Error("Expected 'pool', got", resp.Data[0].Type)
	}
}
