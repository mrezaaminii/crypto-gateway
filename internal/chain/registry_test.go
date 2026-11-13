package chain

import (
	"context"
	"errors"
	"testing"

	"github.com/mrezaaminii/crypto-gateway/internal/domain/asset"
)

type fakeChain struct{}

func (fakeChain) Balance(ctx context.Context, a asset.Asset, address string) (asset.Amount, error) {
	return asset.Amount{}, nil
}
func (fakeChain) EstimateFee(ctx context.Context, req TransferRequest) (Fee, error) {
	return Fee{}, nil
}
func (fakeChain) Transfer(ctx context.Context, req TransferRequest) (string, error) {
	return "", nil
}

func TestRegistry(t *testing.T) {
	r := NewRegistry()
	r.Register(asset.Ethereum, fakeChain{})

	_, err := r.For(asset.Ethereum)
	if err != nil {
		t.Fatalf("expected adapter for ethereum, got error: %v", err)
	}

	_, err = r.For(asset.Bitcoin)
	if !errors.Is(err, ErrUnsupportedNetwork) {
		t.Fatalf("expected ErrUnsupportedNetwork, got: %v", err)
	}
}