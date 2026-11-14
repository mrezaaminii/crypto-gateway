package evm

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mrezaaminii/crypto-gateway/internal/chain"
	"github.com/mrezaaminii/crypto-gateway/internal/domain/asset"
)

var _ chain.Chain = (*Adapter)(nil) 

var (
	ErrClientConnection = errors.New("connection to rpc failed")
	ErrInvalidAddress   = errors.New("invalid address")
	ErrBalanceFailure   = errors.New("failed to get balance")
	ErrTokenUnsupported = errors.New("tokens are not supported yet")
	ErrNotImplemented   = errors.New("not implemented")
)

type Adapter struct {
	client *ethclient.Client
}

func New(ctx context.Context, rpcURL string) (*Adapter, error) {
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrClientConnection, err)
	}
	return &Adapter{client: client}, nil
}

func (a *Adapter) Close() {
	a.client.Close()
}

func (a *Adapter) Balance(ctx context.Context, as asset.Asset, address string) (asset.Amount, error) {
	if !common.IsHexAddress(address) {
		return asset.Amount{}, ErrInvalidAddress
	}

	if !as.IsNative() {
		return asset.Amount{}, ErrTokenUnsupported
	} 

	balance, err := a.client.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return asset.Amount{}, fmt.Errorf("%w: %w", ErrBalanceFailure, err)
	}

	return asset.NewAmount(balance), nil
}

func (a *Adapter) EstimateFee(ctx context.Context, req chain.TransferRequest) (chain.Fee, error) {
	return chain.Fee{}, ErrNotImplemented
}

func (a *Adapter) Transfer(ctx context.Context, req chain.TransferRequest) (string, error) {
	return "", ErrNotImplemented
}