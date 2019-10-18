package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/majordomusio/url-shortener/internal/store"
	"google.golang.org/appengine"
)

// DebugEndpoint maps to GET /
func DebugEndpoint(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)

	ip := c.Query("ip")
	store.CreateGeoLocation(ctx, ip)

	c.JSON(http.StatusOK, "")
}
