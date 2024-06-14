package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	// Create a new Gin engine
	engine = gin.Default()
	// register
	registerHandler()
}

func main() {
	// Run the server on port 8080
	if err := engine.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
