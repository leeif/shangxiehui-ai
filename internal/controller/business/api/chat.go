package api

import (
	"context"
	"encoding/json"
	"io"
	"shangxiehui-ai/internal/common"
	"shangxiehui-ai/internal/facade/dto"
	facadeError "shangxiehui-ai/internal/facade/error"

	"github.com/gin-gonic/gin"
)

func (c *Controller) ChatCompletionTextStream(ctx *gin.Context) *facadeError.Error {

	request := &dto.ChatCompletionTextStreamRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		return facadeError.ErrBadRequest.Wrap(err)
	}

	if len(request.Messages) == 0 {
		return facadeError.ErrBadRequest.Format("messages is empty")
	}

	responseChan, err := c.moonshot.CompletionStream(context.Background(), request.Messages)
	if err != nil {
		return facadeError.ErrServerInternal.Wrap(err)
	}

	ctx.Stream(func(w io.Writer) bool {
		if chunk, ok := <-responseChan; ok {
			event := &dto.ChatTextCompletionStream{
				Role:  common.ChatRoleAssistant,
				Chunk: chunk,
			}

			b, _ := json.Marshal(event)

			ctx.SSEvent("", string(b))
			return true
		}
		return false
	})

	return nil
}
