package types

type (
	// AssetDS is the interal structure used to store assets
	AssetDS struct {
		URI      string `json:"uri,omitempty"`
		URL      string `json:"url" binding:"required"`
		SecretID string `json:"secret_id,omitempty"`
		Cohort   string `json:"cohort,omitempty"`
		// internal metadata
		Created int64 `json:"-"`
	}
)
