package store

import (
	"context"

	"github.com/majordomusio/commons/pkg/util"
	"github.com/majordomusio/platform/pkg/platform"

	"github.com/majordomusio/url-shortener/internal/types"
	"github.com/majordomusio/url-shortener/pkg/api"

	"google.golang.org/appengine/datastore"
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

// AssetKey creates the datastore key for the asset
func AssetKey(ctx context.Context, uri string) *datastore.Key {
	return datastore.NewKey(ctx, api.DatastoreAssets, uri, 0, nil)
}

// CreateAsset stores an asset in the Datastore
func CreateAsset(ctx context.Context, as *api.Asset) error {
	topic := "store.asset.create"

	asset := types.AssetDS{
		URI:      as.URI,
		URL:      as.URL,
		SecretID: as.SecretID,
		Cohort:   as.Cohort,
		Created:  util.Timestamp(),
	}

	key := AssetKey(ctx, as.URI)
	_, err := datastore.Put(ctx, key, &asset)
	if err != nil {
		platform.Error(ctx, topic, err.Error())
		return err
	}

	return nil
}

/*

URI string `json:"uri,omitempty"`
		// URL is the assets real url
		URL string `json:"url" binding:"required"`
		// SecretID can be used to manage the asset
		SecretID string `json:"secret_id,omitempty"`
		// Cohort the asset belongs to
		Cohort string `json:"cohort,omitempty"`

// CreateDefaultModel creates an initial model definition
func CreateDefaultModel(ctx context.Context, clientID string) (*types.ModelDS, error) {

	model := types.ModelDS{
		ClientID: clientID,
		Name:     types.Default,
		Revision: types.DefaultRevision,
		ConfigParams: []types.Parameters{
			{Key: "PythonModule", Value: "model.task"},
			{Key: "RuntimeVersion", Value: "1.12"},
			{Key: "PythonVersion", Value: "2.7"},
		},
		HyperParams: []types.Parameters{
			{Key: "weights", Value: "True"},
			{Key: "latent_factors", Value: "5"},
			{Key: "num_iters", Value: "20"},
			{Key: "regularization", Value: "0.07"},
			{Key: "unobs_weight", Value: "0.01"},
			{Key: "wt_type", Value: "0"},
			{Key: "feature_wt_factor", Value: "130.0"},
			{Key: "feature_wt_exp", Value: "0.08"},
		},
		Events:           []string{types.Default},
		Version:          0,
		TrainingSchedule: 180,
		NextSchedule:     0,
		Created:          util.Timestamp(),
	}

	key := ModelKey(ctx, clientID, types.Default)
	_, err := datastore.Put(ctx, key, &model)
	if err != nil {
		logger.Error(ctx, "backend.model.create", err.Error())
		return nil, err
	}

	return &model, nil
}

// MarkTrained writes an export record back to the datastore with updated metadata
func MarkTrained(ctx context.Context, clientID, name string, trained, next int64) error {

	model, err := GetModel(ctx, clientID, name)
	if err != nil {
		return err
	}

	model.LastTrained = trained
	model.NextSchedule = next

	_, err = datastore.Put(ctx, ModelKey(ctx, clientID, name), model)
	if err != nil {
		return err
	}

	return err
}

*/
