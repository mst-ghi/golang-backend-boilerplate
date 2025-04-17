package bootstrap

import (
	"app/core/config"
	"app/core/engine"
	"app/database"
	"app/internal"
	"app/internal/modules/gateway"
	"app/internal/scheduler"
)

func Serve() {
	config.InitializeAndSet()
	database.InitializeAndConnect()

	internal.Initialize()
	engine.Initialize()

	gateway.Initialize()
	gateway.Serve(engine.GetEngine())

	engine.RegisterRoutes()

	go scheduler.InitScheduler()

	defer gateway.GetSocket().Close()
	engine.Serve()
}
