package fetcher

import (
	"context"
	"fmt"
	"goddra/assetmanager/asset"
)

// LocalFetcher define some asset which has been created in local
type LocalFetcher struct {
	data map[string]*asset.Asset
}

// NewLocalFetcher create a new LocalFetcher
func NewLocalFetcher() *LocalFetcher {
	return &LocalFetcher{
		data: make(map[string]*asset.Asset),
	}
}

// Set create a new asset for a specific name, if their is clash it will override
func (l *LocalFetcher) Set(name string, asset *asset.Asset) {
	l.data[name] = asset
}

// Fetch implements Fetcher.Fetch
func (l *LocalFetcher) Fetch(ctx context.Context, name string) (*asset.Asset, error) {
	if a, ok := l.data[name]; ok {
		return a, nil
	}
	return nil, fmt.Errorf("can't find %s: %w", name, ErrNotFound )
}
