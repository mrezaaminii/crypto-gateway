package evm

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/mrezaaminii/crypto-gateway/internal/domain/asset"
)

const (
	sepoliaNativeTestAddress = "0xfFf9976782d46CC05630D1f6eBAb18b2324d6B14"
	sepoliaUSDC              = "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238"
	sepoliaTokenTestAddress  = "0x8D896d5539B6b5dAAE351d1F679fca818BC0B37C"
)

func newTestAdapter(t *testing.T) (*Adapter, context.Context) {
	t.Helper()
	rpcURL := os.Getenv("SEPOLIA_RPC_URL")
	if rpcURL == "" {
		t.Skip("SEPOLIA_RPC_URL is not set")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	t.Cleanup(cancel)

	adapter, err := New(ctx, rpcURL)

	if err != nil {
		t.Fatalf("New %v: ", err)
	}

	t.Cleanup(adapter.Close)

	return adapter, ctx
}

func TestBalance_Native(t *testing.T) {
	a, ctx := newTestAdapter(t)

	eth := asset.Asset{ID: "ETH", Symbol: "ETH", Network: asset.Ethereum, Decimals: 18, Contract: ""}

	amount, err := a.Balance(ctx, eth, sepoliaNativeTestAddress)

	if err != nil {
		t.Fatalf("Balance: %v", err)
	}

	if amount.String(eth.Decimals) == "0" {
		t.Fatalf("expected a non-zero balance")
	}

	t.Logf("The Amount is %s", amount.String(eth.Decimals))
}

func TestBalance_Token(t *testing.T) {
	a, ctx := newTestAdapter(t)

	usdc := asset.Asset{ID: "USDC_SEPOLIA", Symbol: "USDC", Network: asset.Ethereum, Decimals: 6, Contract: sepoliaUSDC}

	amount, err := a.Balance(ctx, usdc, sepoliaTokenTestAddress)

	if err != nil {
		t.Fatalf("Balance: %v", err)
	}

	if amount.String(usdc.Decimals) == "0" {
		t.Fatalf("expected a non-zero balance")
	}

	t.Logf("The Amount is %s", amount.String(usdc.Decimals))
}
