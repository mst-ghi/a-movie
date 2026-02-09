package swagger

import (
	"app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(router *gin.RouterGroup) {
	docs.SwaggerInfo.Title = "A-Movie"
	docs.SwaggerInfo.Description = "Welcome to the A-Movie API! This service provides a comprehensive interface for accessing a vast database of movies. You can search for specific titles, browse lists with various filters, and retrieve detailed information, including summaries, ratings, and poster images. This documentation will guide you through all the available endpoints and how to use them effectively."
	docs.SwaggerInfo.Version = "0.1.0"

	router.GET(
		"/docs/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(-1),
		),
	)
}
