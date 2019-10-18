package types

import "github.com/majordomusio/url-shortener/pkg/api"

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

	// GeoLocationDS records a IP's geo location
	GeoLocationDS struct {
		IP          string `json:"ip"`
		Host        string `json:"host"`
		ISP         string `json:"isp"`
		City        string `json:"city"`
		CountryCode string `json:"country_code"`
		CountryName string `json:"country_name"`
		Latitude    string `json:"latitude"`
		Longitude   string `json:"longitude"`
	}
)

// AsExternal create an external representation of the asset
func (t *AssetDS) AsExternal() *api.Asset {
	asset := api.Asset{
		URI:       t.URI,
		URL:       t.URL,
		SecretID:  t.SecretID,
		Cohort:    t.Cohort,
		Affiliate: t.Affiliate,
		Tags:      t.Tags,
	}
	return &asset
}
