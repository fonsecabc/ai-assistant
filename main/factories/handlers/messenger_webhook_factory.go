package factories

import (
	"ai-assistant/internal/handlers"
	factories "ai-assistant/main/factories/usecases"
)

func NewMessengerWebhookHandlerFactory() *handlers.MessengerWebhookHandler {
	return handlers.NewMessengerWebhookHandler(
		factories.NewAnswerMessageUseCaseFactory(),
	)
}
