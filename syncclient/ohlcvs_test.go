package syncclient

import "testing"

func TestClient_NetworkPoolOHLCV(t *testing.T) {
	gt := NewClient()
	args := NetworkPoolOHLCVArgs{
		Network:     "eth",
		PoolAddress: "0x60594a405d53811d3bc4766596efd80fd545a270",
		Timeframe:   "day",
	}
	resp, _ := gt.NetworkPoolOHLCV(args)

	if resp.Data.Type != "ohlcv_request_response" {
		t.Error("Expected 'ohlcv_request_response', got", resp.Data.Type)
	}
}
