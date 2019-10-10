package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/majordomusio/commons/pkg/helper"
	"github.com/majordomusio/commons/pkg/util"
	"github.com/majordomusio/url-shortener/internal/store"
	"github.com/majordomusio/url-shortener/pkg/api"
	"google.golang.org/appengine"
)

// DefaultEndpoint maps to GET /
func DefaultEndpoint(c *gin.Context) {
	// TODO: real implementation, logging & auditing
	c.JSON(http.StatusOK, gin.H{"vesion": api.Version, "status": "ok"})
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
	topic := "api.shorten.post"
	ctx := appengine.NewContext(c.Request)

	var asset api.Asset
	err := c.BindJSON(&asset)
	if err == nil {
		uri, _ := util.ShortUUID()
		asset.URI = uri

		err = store.CreateAsset(ctx, &asset)
	}

	helper.StandardJSONResponse(c, topic, asset, err)
}

// RedirectEndpoint receives a URI to be shortened
func RedirectEndpoint(c *gin.Context) {
	//topic := "api.redirect.get"

	short := c.Param("short")

	if short == "" {
		// TODO log this event
		c.String(http.StatusOK, "")
		return
	}

	/*
		a, err := store.Get(short)

		if err != nil {
			// TODO log this event
			c.String(http.StatusOK, "")
			return
		}
	*/
	//helper.StandardJSONResponse(c, topic, a, err)

	// TODO log the event
	//c.Redirect(http.StatusTemporaryRedirect, a.URL)

}
