package query

import (
	"context"
	"errors"
	"testing"

	"github.com/mrezaaminii/crypto-gateway/internal/chain"
	"github.com/mrezaaminii/crypto-gateway/internal/domain/asset"
)

type fakeChain struct {
	balance asset.Amount
}

func (f fakeChain) Balance(ctx context.Context, a asset.Asset, address string) (asset.Amount, error) {
	return f.balance, nil
}

func (f fakeChain) EstimateFee(ctx context.Context, req chain.TransferRequest) (chain.Fee, error) {
	return chain.Fee{}, nil
}

func (f fakeChain) Transfer(ctx context.Context, req chain.TransferRequest) (string, error) {
	return "", nil
}

func TestGetBalance_Handle(t *testing.T) {
	want, err := asset.ParseAmount("1.5", 18)
	if err != nil {
		t.Fatal(err)
	}

	h := NewGetBalanceHandler(newTestCatalog(t), newTestRegistry(want))

	got, err := h.Handle(context.Background(), GetBalance{
		AssetID: "ETH",
		Address: "0xfFf9976782d46CC05630D1f6eBAb18b2324d6B14",
	})
	if err != nil {
		t.Fatalf("Handle: %v", err)
	}

	if got.Amount.String(18) != "1.5" {
		t.Errorf("amount = %s, want 1.5", got.Amount.String(18))
	}
	if got.Asset.ID != "ETH" {
		t.Errorf("asset = %q, want ETH", got.Asset.ID)
	}
}

func TestGetBalance_UnknownAsset(t *testing.T) {
	h := NewGetBalanceHandler(newTestCatalog(t), newTestRegistry(asset.Amount{}))

	_, err := h.Handle(context.Background(), GetBalance{AssetID: "BTC", Address: "0x..."})
	if !errors.Is(err, asset.ErrUnknownAsset) {
		t.Fatalf("err = %v, want ErrUnknownAsset", err)
	}
}

func newTestCatalog(t *testing.T) *asset.Catalog {
	t.Helper()

	cat, err := asset.NewCatalog(asset.Asset{ID: "ETH", Symbol: "ETH", Network: asset.Ethereum, Decimals: 18})

	if err != nil {
		t.Fatal(err)
	}

	return cat
}

func newTestRegistry(balance asset.Amount) *chain.Registry {
	r := chain.NewRegistry()
	r.Register(asset.Ethereum, fakeChain{balance: balance})
	return r
}