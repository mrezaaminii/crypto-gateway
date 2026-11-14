package asset

import "strings"

type Network string

type Asset struct {
	ID       string
	Symbol   string
	Network  Network
	Decimals int
	Contract string
}

const (
	Bitcoin Network = "bitcoin"
	Ethereum Network = "ethereum"
)

func (asset *Asset) IsNative() bool {
	if strings.TrimSpace(asset.Contract) != "" {
		return false
	}

	return true
}