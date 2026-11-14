package evm

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/mrezaaminii/crypto-gateway/internal/domain/asset"
)

func TestBalance_Sepolia(t *testing.T) {
	rpcURL := os.Getenv("SEPOLIA_RPC_URL")
	if rpcURL == "" {
		t.Skip("SEPOLIA_RPC_URL is not set")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	adapter, err := New(ctx, rpcURL)

	if err != nil {
		t.Fatalf("New %v: ", err)
	}

	defer adapter.Close()

	eth := asset.Asset{ID: "ETH", Symbol: "ETH", Network: asset.Ethereum, Decimals: 18, Contract: ""}

	bal, err := adapter.Balance(ctx, eth, "0xfFf9976782d46CC05630D1f6eBAb18b2324d6B14")

	if err != nil {
		t.Fatalf("balance: %v", err)
	}

	t.Logf("balance %s ETH", bal.String(eth.Decimals))
}