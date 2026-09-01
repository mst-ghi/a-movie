package bootstrap

import (
	"app/core/config"
	"app/core/engine"
)

func Serve() {
	config.InitializeAndSet()

	engine.Initialize()

	engine.RegisterRoutes()

	engine.Serve()
}
