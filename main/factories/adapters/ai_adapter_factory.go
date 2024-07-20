package factories

import (
	"ai-assistant/internal/adapters"
	"ai-assistant/main/configs"
)

func NewAIAdapterFactory() *adapters.AIAdapter {
	return adapters.NewAIAdapter(
		configs.Env.OpenAIAPIKey,
		configs.Env.OpenAIModel,
		configs.Env.OpenAIMaxTokens,
	)
}
