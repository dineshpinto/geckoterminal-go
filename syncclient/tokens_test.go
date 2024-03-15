package syncclient

import (
	"testing"
)

func TestClient_NetworkTokenPools(t *testing.T) {
	c := NewClient()
	data, _ := c.NetworkTokenPools("eth", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", 1)
	if data[0].Type != "pool" {
		t.Error("Expected 'pool', got", data[0].Type)
	}
}

func TestClient_NetworkToken(t *testing.T) {
	c := NewClient()
	data, _ := c.NetworkToken("eth", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	if data.Type != "token" {
		t.Error("Expected 'token', got", data.Type)
	}
}

func TestClient_NetworkTokensMultiAddress(t *testing.T) {
	c := NewClient()
	data, _ := c.NetworkTokensMultiAddress("eth", []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"})
	if data[0].Type != "token" {
		t.Error("Expected 'token', got", data[0].Type)
	}
}

func TestClient_NetworkTokenInfo(t *testing.T) {
	c := NewClient()
	data, _ := c.NetworkTokenInfo("eth", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	if data.Type != "token" {
		t.Error("Expected 'token', got", data.Type)
	}
}

func TestClient_NetworkPoolTokenInfo(t *testing.T) {
	c := NewClient()
	data, _ := c.NetworkPoolTokenInfo("eth", "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852")
	if data[0].Type != "token" {
		t.Error("Expected 'token', got", data[0].Type)
	}
}

func TestClient_TokenInfoRecentlyUpdated(t *testing.T) {
	c := NewClient()
	data, _ := c.TokenInfoRecentlyUpdated()
	if data[0].Type != "token" {
		t.Error("Expected 'token', got", data[0].Type)
	}
}
