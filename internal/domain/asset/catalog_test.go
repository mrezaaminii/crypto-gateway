package asset

import (
	"errors"
	"testing"
)

func TestCatalog_Get(t *testing.T) {
	eth := Asset{ID: "ETH", Symbol: "ETH", Network: Ethereum, Decimals: 18}

	c, err := NewCatalog(eth)
	if err != nil {
		t.Fatalf("NewCatalog: %v", err)
	}

	tests := []struct {
		id      string
		wantErr error
	}{
		{"ETH", nil},
		{"BTC", ErrUnknownAsset},
	}

	for _, tt := range tests {
		got, err := c.Get(tt.id)

		if !errors.Is(err, tt.wantErr) {
			t.Errorf("Get(%q) error = %v, want %v", tt.id, err, tt.wantErr)
			continue
		}

		if tt.wantErr == nil && got.ID != tt.id {
			t.Errorf("Get(%q) returned ID %q", tt.id, got.ID)
		}
	}
}

func TestNewCatalog_DuplicateID(t *testing.T) {
	eth := Asset{ID: "ETH", Symbol: "ETH", Network: Ethereum, Decimals: 18}

	_, err := NewCatalog(eth, eth)

	if err == nil {
		t.Fatal("expected an error for duplicate ID, got nil")
	}
}