package store

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/majordomusio/commons/pkg/env"
	"github.com/majordomusio/commons/pkg/util"
	"github.com/majordomusio/platform/pkg/logger"
	"github.com/majordomusio/url-shortener/internal/types"
	"github.com/majordomusio/url-shortener/pkg/api"
)

var dsClient *datastore.Client

func init() {
	ctx := context.Background()
	c, err := datastore.NewClient(ctx, env.Getenv("PROJECT_ID", ""))
	if err != nil {
		logger.Error("DATASTORE: %s", err.Error())
	}
	dsClient = c
}

// Get retrieves the asset
func Get(uri string) (*api.Asset, error) {
	return nil, nil
}

// AssetKey creates the datastore key for an asset
func AssetKey(uri string) *datastore.Key {
	return datastore.NameKey(api.DatastoreAssets, uri, nil)
}

// CreateAsset stores an asset in the Datastore
func CreateAsset(ctx context.Context, as *api.Asset) (string, error) {

	uri, _ := util.ShortUUID()
	secret, _ := util.ShortUUID()

	asset := types.AssetDS{
		URI:      uri,
		URL:      as.URL,
		SecretID: secret,
		Cohort:   as.Cohort,
		Created:  util.Timestamp(),
	}

	k := AssetKey(uri)
	if _, err := dsClient.Put(ctx, k, &asset); err != nil {
		logger.Error("DATASTORE: %s", err.Error())
		return "", err
	}

	return uri, nil
}

// GetAsset retrieves the asset
func GetAsset(ctx context.Context, uri string) (*api.Asset, error) {
	var as types.AssetDS
	k := AssetKey(uri)

	if err := dsClient.Get(ctx, k, &as); err != nil {
		logger.Error("DATASTORE: %s", err.Error())
		return nil, err
	}

	asset := api.Asset{
		URI:      as.URI,
		URL:      as.URL,
		SecretID: as.SecretID,
		Cohort:   as.Cohort,
	}
	return &asset, nil
}
