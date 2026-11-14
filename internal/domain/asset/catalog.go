package asset

import (
	"errors"
	"fmt"
)

var ErrUnknownAsset = errors.New("unknown asset")

type Catalog struct {
	assets map[string]Asset
}

func NewCatalog(assets ...Asset) (*Catalog, error) {
	c := &Catalog{assets: make(map[string]Asset, len(assets))}
	for _, asset := range assets {
		if _, exists := c.assets[asset.ID]; exists {
			return nil, fmt.Errorf("duplicate asset id %q", asset.ID)
		}
		c.assets[asset.ID] = asset
	} 

	return c, nil
}

func (c *Catalog) Get(id string) (Asset, error) {
	asset, ok := c.assets[id]

	if !ok {
		return Asset{}, fmt.Errorf("%w: %q", ErrUnknownAsset, id)
	}

	return asset, nil
}