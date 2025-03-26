package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the router with gin's default settings
	r := gin.Default()

	// Define the routes
	r.GET("/healthy", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "success"}) })

	// Open the server on port :8000
	r.Run(":8000")
}
