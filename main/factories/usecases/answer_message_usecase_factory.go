package factories

import (
	"ai-assistant/domain/usecases"
	factories "ai-assistant/main/factories/adapters"
)

func NewAnswerMessageUseCaseFactory() *usecases.AnswerMessageUseCase {
	return usecases.NewAnswerMessageUseCase(
		factories.NewAIAdapterFactory(),
		factories.NewMessengerAdapterFactory(),
	)
}
