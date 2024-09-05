package route

import (
	"net/http"
	"shangxiehui-ai/internal/facade/dto"
	"shangxiehui-ai/internal/utils/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	facadeError "shangxiehui-ai/internal/facade/error"
)

// only used in this package
var (
	internalLogger *logger.KiwiLogger
)

func HandleErrorStream(f func(*gin.Context) *facadeError.Error) func(*gin.Context) {

	return func(c *gin.Context) {

		var facadeErr *facadeError.Error

		// execute controller method
		if facadeErr = f(c); facadeErr != nil {
			responseError(c, facadeErr)
			return
		}
	}
}

func HandleError[T any](f func(*gin.Context) (T, *facadeError.Error)) func(*gin.Context) {

	return func(c *gin.Context) {

		data, err := f(c)
		if err != nil {
			responseError(c, err)
			return
		}

		c.JSON(http.StatusOK, &dto.Response{
			Status: "ok",
			Data:   data,
		})
	}
}

func MiddlewareWrapper(f func(*gin.Context) *facadeError.Error) func(*gin.Context) {

	return func(c *gin.Context) {

		if err := f(c); err != nil {
			responseError(c, err)
			c.Abort()
			return
		}

		c.Next()
	}
}

func SkipMiddlewareError(f func(*gin.Context) *facadeError.Error) func(*gin.Context) {

	return func(c *gin.Context) {

		if err := f(c); err != nil {
			internalLogger.Warn("Skip middleware error: ", zap.Error(err))
		}

		c.Next()
	}
}

func responseError(c *gin.Context, err *facadeError.Error) {
	response := &dto.Response{}
	path := c.Request.URL.Path
	internalLogger.Error("Request failed", zap.String("path", path), zap.Error(err))
	response.Status = "error"
	response.Error = err
	c.JSON(err.StatusCode(), response)
}
