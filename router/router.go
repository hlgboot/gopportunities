package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the router with gin's default settings
	r := gin.Default()

	// Initialize routes
	initializeRoutes(r)

	// Open the server on port :8000
	r.Run(":8000")
}
