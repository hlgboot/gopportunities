package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {

	// Define the routes
	v1 := r.Group("/api/v1")
	{
		// Verify api availability
		v1.GET("/healthy", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"success": true}) })

		//
		v1.GET("/opening", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "GET Opening"}) })

		//
		v1.POST("/opening", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "POST Opening"}) })

		//
		v1.DELETE("/opening", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "DELETE Opening"}) })

		//
		v1.PUT("/opening", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "PUT Opening"}) })

		//
		v1.GET("/openings", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "GET Openings"}) })
	}
}
