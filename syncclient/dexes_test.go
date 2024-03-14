package syncclient

import "testing"

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
