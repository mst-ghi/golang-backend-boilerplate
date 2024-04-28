package bootstrap

import (
	"app/core"
	"app/core/config"
	"app/core/engine"
	"app/core/scheduler"
	"app/database"
	"app/domain/gateway"
)

func Serve() {
	config.InitializeAndSet()
	database.InitializeAndConnect()

	core.Initialize()
	engine.Initialize()

	gateway.Initialize()
	gateway.Serve(engine.GetEngine())

	engine.RegisterRoutes()

	go scheduler.InitScheduler()

	defer gateway.GetSocket().Close()
	engine.Serve()
}
