package domain

import (
	"app/core"
	"app/domain/movies"

	"github.com/gin-gonic/gin"
)

// @tags    App
// @router	/api [get]
// @summary	app route, get healthy status
func HomeRoute(c *gin.Context) {
	ctrl := core.GetController()
	ctrl.Success(c, map[string]string{
		"healthy": "I'm OK...",
	})
}

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("", HomeRoute)

	v1Group := router.Group("/v1")
	{
		movies.RegisterRoutes(v1Group)
	}
}
