package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/majordomusio/commons/pkg/util"
	"github.com/majordomusio/url-shortener/internal/store"
	"github.com/majordomusio/url-shortener/internal/types"
	"github.com/majordomusio/url-shortener/pkg/api"
	"google.golang.org/appengine"
)

// DefaultEndpoint maps to GET /
func DefaultEndpoint(c *gin.Context) {
	// TODO: real implementation, logging & auditing
	c.JSON(http.StatusOK, gin.H{"app": api.FullName, "vesion": api.Version, "status": "ok"})
}

// RobotsEndpoint maps to GET /robots.txt
func RobotsEndpoint(c *gin.Context) {
	// simply write text back ...
	c.Header("Content-Type", "text/plain")

	// a simple robots.txt file, disallow the API
	c.Writer.Write([]byte("User-agent: *\n\n"))
	c.Writer.Write([]byte("Disallow: /api/\n"))
}

// ShortenEndpoint receives a URI to be shortened
func ShortenEndpoint(c *gin.Context) {
	//topic := "api.shorten.post"
	ctx := appengine.NewContext(c.Request)

	var asset api.Asset
	var uri string

	err := c.BindJSON(&asset)
	if err == nil {
		uri, err = store.CreateAsset(ctx, &asset)
		asset.URI = uri
	}

	standardJSONResponse(c, asset, err)
}

// RedirectEndpoint receives a URI to be shortened
func RedirectEndpoint(c *gin.Context) {
	//topic := "api.redirect.get"
	ctx := appengine.NewContext(c.Request)

	uri := c.Param("uri")
	if uri == "" {
		// TODO log this event
		c.String(http.StatusOK, "42")
		return
	}

	a, err := store.GetAsset(ctx, uri)
	if err != nil {
		// TODO log this event
		c.String(http.StatusOK, "42")
		return
	}

	// audit, i.e. extract some user data
	m := types.MeasurementDS{
		URI:            uri,
		User:           "anonymous",
		IP:             c.ClientIP(),
		UserAgent:      strings.ToLower(c.GetHeader("User-Agent")),
		AcceptLanguage: strings.ToLower(c.GetHeader("Accept-Language")),
		Created:        util.Timestamp(),
	}
	store.CreateMeasurement(ctx, &m)

	// TODO log the event
	c.Redirect(http.StatusTemporaryRedirect, a.URL)
}
