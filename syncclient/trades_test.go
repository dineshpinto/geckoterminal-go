package syncclient

import (
	"testing"
)

func TestClient_NetworkPoolTrades(t *testing.T) {
	gt := NewClient()
	resp, _ := gt.NetworkPoolTrades("eth", "0x60594a405d53811d3bc4766596efd80fd545a270", 1)

	if resp.Data[0].Type != "trade" {
		t.Error("Expected 'trade', got", resp.Data[0].Type)
	}
}
