package syncclient

import "testing"

func TestClient_NetworkPoolOHLCV(t *testing.T) {
	gt := GeckoTerminalClient()
	args := NetworkPoolOHLCVArgs{
		Network:     "eth",
		PoolAddress: "0x60594a405d53811d3bc4766596efd80fd545a270",
		Timeframe:   "day",
	}
	data, _ := gt.NetworkPoolOHLCV(args)

	if data.Type != "ohlcv_request_response" {
		t.Error("Expected 'ohlcv_request_response', got", data.Type)
	}
}
