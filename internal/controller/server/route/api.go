package route

import (
	"github.com/gin-gonic/gin"
)

// RegisterApiV1 register route for api server
func (route *Route) RegisterApiV1(gin *gin.Engine) {
	gin.GET("/ping", HandleError(route.apiController.Ping))

	v1 := gin.Group("/v1")

	chat := v1.Group("/chat")
	{
		chat.POST(
			"/completion/text/stream",
			HandleErrorStream(route.apiController.ChatCompletionTextStream))
	}
}
