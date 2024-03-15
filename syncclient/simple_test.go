package syncclient

import "testing"

func TestClient_NetworkAddressesTokenPrice(t *testing.T) {
	gt := NewClient()
	addresses := []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"}
	resp, _ := gt.NetworkAddressesTokenPrice("eth", addresses)

	if resp.Data.Type != "simple_token_price" {
		t.Error("Expected 'simple_token_price', got", resp.Data.Type)
	}
}
