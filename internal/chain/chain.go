package chain

import (
	"context"

	"github.com/mrezaaminii/crypto-gateway/internal/domain/asset"
)

type TransferRequest struct {
	Asset asset.Asset
	From string 
	To string 
	Amount asset.Amount
	Memo string
}

type Fee struct {
	Asset asset.Asset
	Amount asset.Amount
}

type Chain interface {
	Balance(ctx context.Context, a asset.Asset, address string) (asset.Amount, error)
	EstimateFee(ctx context.Context, req TransferRequest) (Fee, error)
	Transfer(ctx context.Context, req TransferRequest) (txHash string, err error)
}