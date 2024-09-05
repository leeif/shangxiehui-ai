package server

import (
	"shangxiehui-ai/internal/utils/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ginlogger = func(logger *logger.KiwiLogger) gin.HandlerFunc {
		return func(c *gin.Context) {
			start := time.Now()
			// some evil middlewares modify this values
			path := c.Request.URL.Path
			c.Next()

			end := time.Now()
			latency := end.Sub(start)

			fields := []zapcore.Field{
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("ip", c.ClientIP()),
				zap.String("useragent", c.Request.UserAgent()),
				zap.Duration("latency", latency),
			}

			if len(c.Errors) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range c.Errors.Errors() {
					logger.Error(e, fields...)
				}
			} else {
				logger.Info(path, fields...)
			}
		}
	}
)
