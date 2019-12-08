package assetmanager

import (
	"context"
	"goddra/assetmanager/asset"
	"goddra/assetmanager/fetcher"
	"sort"
)

type fetcherRank struct {
	Prio int
	Fetcher fetcher.Fetcher
}

// AssetManager add assets
type AssetManager struct {
	fetchers []fetcherRank
	cache    map[string]*asset.Asset
}

func (am *AssetManager) Fetch(ctx context.Context, name string) (*asset.Asset, error) {
	var a *asset.Asset
	var err error
	for _, f := range am.fetchers {
		a, err = f.Fetcher.Fetch(ctx, name)
		if err == nil {
			return a, nil
		}
	}
	return nil, err
}

// AddFetchers will add way to fetch asset
func (am *AssetManager) AddFetcher(prio int, fetcher fetcher.Fetcher) *AssetManager {
	am.fetchers = append(am.fetchers, fetcherRank{
		Prio:    prio,
		Fetcher: fetcher,
	})
	sort.Slice(am.fetchers, func(i, j int) bool {
		return am.fetchers[i].Prio < am.fetchers[j].Prio
	})
	return am
}
