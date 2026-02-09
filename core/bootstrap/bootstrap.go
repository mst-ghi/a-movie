package bootstrap

import (
	"app/core"
	"app/core/config"
	"app/core/engine"
)

func Serve() {
	config.InitializeAndSet()

	core.Initialize()
	engine.Initialize()

	engine.RegisterRoutes()

	engine.Serve()
}
