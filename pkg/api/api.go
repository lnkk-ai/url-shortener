package api

const (
	// Version is the human readable version string of this build
	Version string = "1.0"

	// APIPrefix is the common namespace prefix of API releated routes
	APIPrefix string = "/api/1"

	// RedirectPrefix is used to build the redirect URLs
	RedirectPrefix string = "/r"
)

type (
	// Asset is the basic entity used in the shortener
	Asset struct {
		URI      string `json:"uri,omitempty"`
		URL      string `json:"url" binding:"required"`
		SecretID string `json:"secret_id,omitempty"`
	}
)
