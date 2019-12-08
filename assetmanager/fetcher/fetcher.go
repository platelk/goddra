package fetcher

import (
	"context"
	"fmt"
	"goddra/assetmanager/asset"
)

var ErrNotFound = fmt.Errorf("asset is not found")

// Fetcher is the interface which define how to fetch an [Asset]
type Fetcher interface {
	// Fetch an asset base on a name
	Fetch(ctx context.Context, name string) (*asset.Asset, error)
}
