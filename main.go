package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Server is starting...")

	// Set Gin mode, define basic server, and make server use cors
	gin.SetMode(gin.ReleaseMode)
	httpServer := gin.Default()
	httpServer.Use(cors.Default())
	httpServer.Use(static.Serve("/", static.LocalFile("./dist", true)))

	// Define an http client to be used by all necessary handlers
	client := &http.Client{
		Timeout: time.Second * 20,
		Transport: &http.Transport{
			MaxIdleConns:        20,
			MaxIdleConnsPerHost: 5,
		},
	}

	// Load the necessary routes into the gin server
	httpServer.GET("/choices", ChoicesResp)
	httpServer.GET("/choice", ChoiceResp(client))
	httpServer.POST("/play", PlayResp(client))

	// Log the gin server as its been defined and have it run locally
	log.WithField("server", httpServer).Info("Default Gin Server Created.")
	httpServer.Run("0.0.0.0:80")
}
