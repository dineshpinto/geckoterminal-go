package syncclient

import (
	"testing"
)

func TestClient_NetworkPoolTrades(t *testing.T) {
	gt := NewClient()
	data, _ := gt.NetworkPoolTrades("eth", "0x60594a405d53811d3bc4766596efd80fd545a270", 1)

	if data[0].Type != "trade" {
		t.Error("Expected 'trade', got", data[0].Type)
	}
}
