package route

import (
	"shangxiehui-ai/internal/controller/business/api"
	"shangxiehui-ai/internal/controller/server/middleware"
	"shangxiehui-ai/internal/utils/logger"
)

type Route struct {
	apiController *api.Controller
	mw            *middleware.Middleware
	logger        *logger.KiwiLogger
}

func NewRoute(apiController *api.Controller,
	mw *middleware.Middleware,
	logger *logger.KiwiLogger) *Route {

	internalLogger = logger

	return &Route{
		apiController: apiController,
		mw:            mw,
		logger:        logger,
	}
}
