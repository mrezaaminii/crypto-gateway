package evm

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
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
	ErrNotImplemented   = errors.New("not implemented")
	ErrNotValidContract = errors.New("not valid contract")
	ErrCallContract     = errors.New("calling contract failed")
	ErrMismatchContract = errors.New("mismatch in getting contract")
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

	owner := common.HexToAddress(address)

	if as.IsNative() {
		return a.nativeBalance(ctx, owner)
	}

	return a.tokenBalance(ctx, as, owner)
}

func (a *Adapter) EstimateFee(ctx context.Context, req chain.TransferRequest) (chain.Fee, error) {
	return chain.Fee{}, ErrNotImplemented
}

func (a *Adapter) Transfer(ctx context.Context, req chain.TransferRequest) (string, error) {
	return "", ErrNotImplemented
}

func (a *Adapter) nativeBalance(ctx context.Context, owner common.Address) (asset.Amount, error) {
	balance, err := a.client.BalanceAt(ctx, owner, nil)
	if err != nil {
		return asset.Amount{}, fmt.Errorf("%w: %w", ErrBalanceFailure, err)
	}

	return asset.NewAmount(balance), nil
}

func (a *Adapter) tokenBalance(ctx context.Context, as asset.Asset, owner common.Address) (asset.Amount, error) {
	if !common.IsHexAddress(as.Contract) {
		return asset.Amount{}, ErrNotValidContract
	}

	selector := []byte{0x70, 0xa0, 0x82, 0x31}

	data := append(selector, common.LeftPadBytes(owner.Bytes(), 32)...)

	contract := common.HexToAddress(as.Contract)

	out, err := a.client.CallContract(ctx, ethereum.CallMsg{To: &contract, Data: data}, nil)

	if err != nil {
		return asset.Amount{}, fmt.Errorf("%w: %w", ErrCallContract, err)
	}

	if len(out) != 32 {
		return asset.Amount{}, ErrMismatchContract
	}

	return asset.NewAmount(new(big.Int).SetBytes(out)), nil
}
