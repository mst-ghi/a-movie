package engine

import (
	"log"
	"net/http"

	"app/core/config"
	"app/core/middlewares"
	"app/core/swagger"
	"app/domain"
	"app/pkg/handlers"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func GetEngine() *gin.Engine {
	return engine
}

func Initialize() {
	gin.SetMode(config.Get("GIN_MODE"))

	engine = gin.Default()
	engine.SetTrustedProxies(nil)
	engine.RedirectTrailingSlash = true
	engine.RedirectFixedPath = true

	engine.LoadHTMLGlob("templates/*")
	engine.Static("/assets", "assets")

	engine.Use(middlewares.Cors())
	engine.Use(gin.CustomRecovery(handlers.InternalErrorHandler))

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
}

func Serve(addr ...string) {
	runAddress := config.GetRunAddress()

	if addr != nil {
		runAddress = addr[0]
	}

	log.Printf("server listening on %s", runAddress)
	if err := engine.Run(runAddress); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}

func RegisterRoutes() {
	routerGroup := GetEngine().Group("api")

	domain.RegisterRoutes(routerGroup)

	swagger.RegisterSwagger(routerGroup)
}
