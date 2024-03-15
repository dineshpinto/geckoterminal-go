package syncclient

import (
	"testing"
)

func TestClient_Networks(t *testing.T) {
	gt := NewClient()
	resp, _ := gt.Networks(1)

	if resp.Data[0].Type != "network" {
		t.Error("Expected 'network', got", resp.Data[0].Type)
	}
	if resp.Data[0].Id != "eth" {
		t.Error("Expected 'eth', got", resp.Data[0].Id)
	}
}
