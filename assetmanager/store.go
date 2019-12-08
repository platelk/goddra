package assetmanager

import "goddra/assetmanager/asset"

// Store interface define how to store a retrieve asset
type Store interface {
	Set(name string, asset *asset.Asset)
}
