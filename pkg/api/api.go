package api

const (
	// FullName is the name of the service
	FullName string = "shadowman-the-bot (shtb), url-shortener"

	// Version is the human readable version string of this build
	Version string = "1.0"

	// APIPrefix is the common namespace prefix of API releated routes
	APIPrefix string = "/api/1"

	// RedirectPrefix is used to build the redirect URLs
	RedirectPrefix string = "/r"

	// DatastoreAssets collection ASSETS
	DatastoreAssets string = "ASSETS"
	// DatastoreMeasurement collection MEASUREMENT
	DatastoreMeasurement string = "MEASUREMENT"
)

type (
	// Asset is the basic entity used in the shortener
	Asset struct {
		// URI identifies the asset
		URI string `json:"uri,omitempty"`
		// URL is the assets real url
		URL string `json:"url" binding:"required"`
		// SecretID can be used to manage the asset
		SecretID string `json:"secret_id,omitempty"`
		// Cohort the asset belongs to
		Cohort string `json:"cohort,omitempty"`
	}
)

/*

	// Parameters is a generic struct to store configuration parameters
	Parameters struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

*/
