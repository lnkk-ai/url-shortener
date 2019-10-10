package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/majordomusio/commons/pkg/helper"
	"github.com/majordomusio/commons/pkg/util"
	"github.com/shadowman-the-bot/shtb-url-shortener/pkg/api"
	"github.com/shadowman-the-bot/shtb-url-shortener/pkg/store"
)

// ShortenEndpoint receives a URI to be shortened
func ShortenEndpoint(c *gin.Context) {
	var asset api.Asset
	topic := "api.shorten.post"

	err := c.BindJSON(&asset)
	if err == nil {
		uri, _ := util.ShortUUID()
		asset.URI = uri

		store.Create(&asset)
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

	a, err := store.Get(short)

	if err != nil {
		// TODO log this event
		c.String(http.StatusOK, "")
		return
	}

	//helper.StandardJSONResponse(c, topic, a, err)

	// TODO log the event
	c.Redirect(http.StatusTemporaryRedirect, a.URL)

}
