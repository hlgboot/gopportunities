package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hlgboot/gopportunities/handler"

	docs "github.com/hlgboot/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {

	// Initialize Handler
	handler.InitializeHandler()

	// Swagger Docs Config
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// Define the routes
	v1 := r.Group(basePath)
	{
		// Verify api availability
		v1.GET("/healthy", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"success": true}) })

		// SHOW Opening
		v1.GET("/opening", handler.ShowOpeningHandler)

		// CREATE Opening
		v1.POST("/opening", handler.CreateOpeningHandler)

		// DELETE Opening
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		// UPDATE Opening
		v1.PUT("/opening", handler.UpdateOpeningHandler)

		// LIST Openings
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	// Initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
