package query

import (
	"context"

	"github.com/mrezaaminii/crypto-gateway/internal/chain"
	"github.com/mrezaaminii/crypto-gateway/internal/domain/asset"
)

type GetBalance struct {
	AssetID string
	Address string
}

type BalanceResult struct {
	Asset asset.Asset
	Amount asset.Amount
	Address string
}

type GetBalanceHandler struct {
	assets *asset.Catalog
	chains *chain.Registry
}

func NewGetBalanceHandler(assets *asset.Catalog, chains *chain.Registry) *GetBalanceHandler {
	return &GetBalanceHandler{assets: assets, chains: chains}
}

func (h *GetBalanceHandler) Handle(ctx context.Context, q GetBalance) (BalanceResult, error) {
	a, err := h.assets.Get(q.AssetID)

	if err != nil {
		return BalanceResult{}, err
	}
	c, err := h.chains.For(a.Network)

	if err != nil {
		return BalanceResult{}, err
	}

	amount, err := c.Balance(ctx, a, q.Address)

	if err != nil {
		return BalanceResult{}, err
	}

	return BalanceResult{Asset: a, Amount: amount, Address: q.Address}, nil
}