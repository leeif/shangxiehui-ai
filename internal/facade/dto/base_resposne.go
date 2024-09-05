package dto

import (
	"shangxiehui-ai/internal/facade/error"
)

// Response godoc
// @Description Base response.
type Response struct {
	Status string       `json:"status" swaggertype:"string"`
	Error  *error.Error `json:"error"`
	Data   interface{}  `json:"data" swaggertype:"object"`
}
