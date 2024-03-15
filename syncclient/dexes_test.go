package syncclient

import "testing"

func TestClient_NetworkDexes(t *testing.T) {
	gt := NewClient()
	resp, _ := gt.NetworkDexes("eth", 1)

	if resp.Data[0].Type != "dex" {
		t.Error("Expected 'dex', got", resp.Data[0].Type)
	}
	if resp.Data[0].Id != "uniswap_v2" {
		t.Error("Expected 'uniswap_v2', got", resp.Data[0].Id)
	}
}
