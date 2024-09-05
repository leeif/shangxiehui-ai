package infrastructure

import (
	"shangxiehui-ai/internal/infrastructure/llm/moonshot"
	"shangxiehui-ai/pkg/cache"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	moonshot.NewClient,
	cache.NewMemCache,
)
