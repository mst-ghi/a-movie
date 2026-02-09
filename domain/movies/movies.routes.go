package movies

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewMoviesController()

	authGroup := router.Group("/movies")
	{
		authGroup.GET("", ctrl.FindAll)
		authGroup.GET("/search", ctrl.Search)
	}
}
