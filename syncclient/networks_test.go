package syncclient

import (
	"testing"
)

func TestClient_Networks(t *testing.T) {
	gt := NewClient()
	data, _ := gt.Networks(1)

	if data[0].Type != "network" {
		t.Error("Expected 'network', got", data[0].Type)
	}
	if data[0].Id != "eth" {
		t.Error("Expected 'eth', got", data[0].Id)
	}
}
