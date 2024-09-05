package controller

import (
	"shangxiehui-ai/internal/controller/business/api"
	"shangxiehui-ai/internal/controller/server"
	"shangxiehui-ai/internal/controller/server/middleware"
	"shangxiehui-ai/internal/controller/server/route"
	"shangxiehui-ai/internal/utils/logger"

	"go.uber.org/fx"
)

var ServerModule = fx.Provide(
	route.NewRoute,
	api.NewController,
	middleware.NewMiddleware,
	server.NewAPIServer,

	// logger
	fx.Annotate(
		logger.NewKiwiLogger,
		fx.OnStop(func(logger *logger.KiwiLogger) {
			_ = logger.Sync()
		}),
	),
)
