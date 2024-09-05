package api

import (
	"shangxiehui-ai/internal/facade/dto"
	facadeError "shangxiehui-ai/internal/facade/error"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Ping(_ *gin.Context) (*dto.OperationResponse, *facadeError.Error) {
	return &dto.OperationResponse{
		Success: true,
	}, nil
}
