package store

import (
	"github.com/majordomusio/url-shortener/pkg/api"
)

var cache map[string]*api.Asset

func init() {
	cache = make(map[string]*api.Asset)
}

// Create stores the asset
func Create(asset *api.Asset) error {
	cache[asset.URI] = asset
	return nil
}

// Get retrieves the asset
func Get(uri string) (*api.Asset, error) {
	a, ok := cache[uri]
	if ok {
		return a, nil
	}
	return nil, nil
}
