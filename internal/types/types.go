package types

type (
	// AssetDS is the interal structure used to store assets
	AssetDS struct {
		URI string `json:"uri,omitempty"`
		URL string `json:"url" binding:"required"`
		// ownership etc
		Owner    string `json:"owner,omitempty"`
		SecretID string `json:"secret_id,omitempty"`
		// segmentation
		Source    string `json:"source,omitempty"`
		Cohort    string `json:"cohort,omitempty"`
		Affiliate string `json:"affiliate,omitempty"`
		Tags      string `json:"tags,omitempty"`
		// internal metadata
		Created int64 `json:"-"`
	}

	// MeasurementDS records events
	MeasurementDS struct {
		URI            string `json:"uri" binding:"required"`
		User           string `json:"user" binding:"required"`
		IP             string `json:"ip,omitempty"`
		UserAgent      string `json:"user_agent,omitempty"`
		AcceptLanguage string `json:"accept_language,omitempty"`
		// internal metadata
		Created int64 `json:"-"`
	}
)
