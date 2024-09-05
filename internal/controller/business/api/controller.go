package api

import (
	"shangxiehui-ai/config"
	"shangxiehui-ai/internal/infrastructure/llm/moonshot"
)

type Controller struct {
	config   *config.Config
	moonshot *moonshot.Client
}

func NewController(
	config *config.Config,
	moonshot *moonshot.Client,
) (*Controller, error) {
	return &Controller{
		config:   config,
		moonshot: moonshot,
	}, nil
}
