package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerHandler() {
	// Define a route for the "/" path
	engine.GET("/alfafrens_frame/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pang",
		})
	})

	engine.GET("/alfafrens_frame/channel_name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pang",
		})
	})
}
