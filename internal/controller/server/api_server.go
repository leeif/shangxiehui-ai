package server

import (
	"context"
	"errors"
	"net/http"
	"shangxiehui-ai/config"
	"shangxiehui-ai/internal/controller/server/route"
	"shangxiehui-ai/internal/utils/logger"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type APIServer struct {
	Engine *gin.Engine
}

func NewAPIServer(lc fx.Lifecycle, logger *logger.KiwiLogger, cfg *config.Config, route *route.Route) *APIServer {
	logger = logger.With(zap.Namespace("apiserver"))

	engine := gin.New()
	engine.Use(ginlogger(logger))
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"},
		AllowHeaders: []string{"*"},
		MaxAge:       12 * time.Hour,
	}))
	ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	srv := &http.Server{
		Addr:    ":" + cfg.APIServer.Port.String(),
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("API server start listening", zap.String("port", cfg.APIServer.Port.String()))
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error("API server listen error", zap.Error(err))
				}
			}()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Info("API server shutdown...")
			if err := srv.Shutdown(ctx); err != nil {
				logger.Error("API server shutdown error", zap.Error(err))
			}
			return nil
		},
	})

	return &APIServer{
		Engine: engine,
	}
}
