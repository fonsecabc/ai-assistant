package factories

import (
	"ai-assistant/internal/adapters"
	"ai-assistant/main/configs"
)

func NewMessengerAdapterFactory() *adapters.MessengerAdapter {
	return adapters.NewMessengerAdapter(
		configs.Env.HubspotAccessToken,
		configs.Env.HubspotAPIBaseURL,
	)
}
