package fetcher

import (
	"context"
	"fmt"
	"goddra/assetmanager/asset"
	"net/http"
)

// JsFetcher is an implementation for browser
type JsFetcher struct {
	basePath string
}

// NewJsFetcher will instantiate an JsFetcher
func NewJsFetcher(basePath string) *JsFetcher {
	return &JsFetcher{
		basePath:basePath,
	}
}

// Fetch implements Fetcher.Fetch
func (j *JsFetcher) Fetch(ctx context.Context, name string) (*asset.Asset, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s", name), nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Add("js.fetch:mode", "cors")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("wrong status code, received %d", resp.StatusCode)
	}
	return asset.NewAssetFromReader(name, resp.Body), nil
}
