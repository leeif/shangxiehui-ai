package middleware

import (
	"shangxiehui-ai/config"
	"shangxiehui-ai/internal/utils/logger"
	"shangxiehui-ai/pkg/cache"

	"go.uber.org/zap"
)

type Middleware struct {
	memCache *cache.MemCache
	logger   *logger.KiwiLogger
	cfg      *config.Config
}

func NewMiddleware(
	cfg *config.Config,
	memCache *cache.MemCache,
	logger *logger.KiwiLogger) (*Middleware, error) {
	return &Middleware{
		cfg:      cfg,
		memCache: memCache,
		logger:   logger.With(zap.Namespace("middleware")),
	}, nil
}
