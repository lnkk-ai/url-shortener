package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	a "github.com/lnkk-ai/url-shortener/internal/api"
	"github.com/lnkk-ai/url-shortener/internal/store"
	"github.com/lnkk-ai/url-shortener/pkg/api"
	"github.com/lnkk-ai/url-shortener/pkg/errorreporting"
)

func main() {
	// setup shutdown handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		shutdown()
		os.Exit(1)
	}()

	// basic config
	gin.DisableConsoleColor()
	// a new router
	router := gin.New()
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// default endpoints that are not part of the API namespace
	router.GET("/", a.DefaultEndpoint)
	router.GET("/robots.txt", a.RobotsEndpoint)
	//router.GET("/debug", a.DebugEndpoint)

	// group by namespace
	router.POST(api.APIPrefix+"/short", a.ShortenEndpoint)
	router.GET(api.RedirectPrefix+"/:uri", a.RedirectEndpoint)

	// start the router on port 8080, unless ENV PORT is set to something else
	router.Run()
}

func shutdown() {
	log.Printf("Shutting down ...")

	store.Close()
	errorreporting.Close()

	log.Printf("... done")
}
