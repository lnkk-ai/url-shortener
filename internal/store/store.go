package store

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/majordomusio/commons/pkg/env"
	"github.com/majordomusio/commons/pkg/util"
	"github.com/majordomusio/url-shortener/internal/types"
	"github.com/majordomusio/url-shortener/pkg/api"
	"github.com/majordomusio/url-shortener/pkg/errorreporting"
)

var dsClient *datastore.Client

func init() {
	ctx := context.Background()
	c, err := datastore.NewClient(ctx, env.Getenv("PROJECT_ID", ""))
	if err != nil {
		errorreporting.Report(err)
	}
	dsClient = c
}

// Close does the clean-up
func Close() {
	dsClient.Close()
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
		URI:       uri,
		URL:       as.URL,
		Owner:     "anonymous",
		SecretID:  secret,
		Source:    valueWithDefault(as.Source, "redhat-capgemini.slack.com"),
		Cohort:    as.Cohort,
		Affiliate: as.Affiliate,
		Tags:      as.Tags,
		Created:   util.Timestamp(),
	}

	k := AssetKey(uri)
	if _, err := dsClient.Put(ctx, k, &asset); err != nil {
		errorreporting.Report(err)
		return "", err
	}

	return uri, nil
}

// GetAsset retrieves the asset
func GetAsset(ctx context.Context, uri string) (*api.Asset, error) {
	var as types.AssetDS
	k := AssetKey(uri)

	if err := dsClient.Get(ctx, k, &as); err != nil {
		errorreporting.Report(err)
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

// CreateMeasurement records a link activation
func CreateMeasurement(ctx context.Context, m *types.MeasurementDS) error {

	// anonimize the IP to be GDPR compliant
	m.IP = anonimize(m.IP)

	k := datastore.IncompleteKey(api.DatastoreMeasurement, nil)
	if _, err := dsClient.Put(ctx, k, m); err != nil {
		errorreporting.Report(err)
		return err
	}

	return nil
}

// Anonimize the IP to be GDPR compliant
func anonimize(ip string) string {
	if strings.ContainsRune(ip, 58) {
		// IPv6
		parts := strings.Split(ip, ":")
		return fmt.Sprintf("%s:%s:%s:0000:0000:0000:0000:0000", parts[0], parts[1], parts[2])
	}
	// IPv4
	parts := strings.Split(ip, ".")
	return fmt.Sprintf("%s.%s.%s.0", parts[0], parts[1], parts[2])
}

// valueWithDefault returns the value of a default if empty
func valueWithDefault(value, def string) string {
	if value == "" {
		return value
	}
	return def
}
