package chain

import (
	"errors"
	"fmt"

	"github.com/mrezaaminii/crypto-gateway/internal/domain/asset"
)

var ErrUnsupportedNetwork = errors.New("Unsupported network!")

type Registry struct {
	chains map[asset.Network]Chain
}

func NewRegistry() *Registry {
	return &Registry{chains: map[asset.Network]Chain{}}
}

func (r *Registry) Register(n asset.Network, c Chain) {
	r.chains[n] = c
}

func (r *Registry) For(n asset.Network) (Chain, error) {
	chain, ok := r.chains[n]
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnsupportedNetwork, n)
	}

	return chain ,nil
}