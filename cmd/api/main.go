package main

import (
	"html/template"
	"net/http"
	"twelveGo/config"
	"twelveGo/pkg/api"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	router := gin.Default()

	// Load HTML templates
	router.SetHTMLTemplate(template.Must(template.ParseFiles("web/templates/index.html")))

	// Serve index.html at the root route
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Port":     cfg.Port,
			"DBURL":    cfg.DBURL,
			"MongoURL": cfg.MongoURL,
		})
	})

	router.GET("/heartbeat", api.Heartbeat)
	router.GET("/memberheartbeat", api.MemberHeartbeat)

	// Start the server
	router.Run(cfg.Port)
}

