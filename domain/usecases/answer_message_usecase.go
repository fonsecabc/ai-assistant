package usecases

import (
	"ai-assistant/internal/adapters"
	"fmt"
)

type AnswerMessageUseCase struct {
	AIAdapter       *adapters.AIAdapter
	MessagerAdapter *adapters.MessengerAdapter
}

func NewAnswerMessageUseCase(aiAdapter *adapters.AIAdapter, messagerAdapter *adapters.MessengerAdapter) *AnswerMessageUseCase {
	return &AnswerMessageUseCase{
		AIAdapter:       aiAdapter,
		MessagerAdapter: messagerAdapter,
	}
}

func (a *AnswerMessageUseCase) Perform(messageId, threadId string) error {
	m, err := a.MessagerAdapter.GetMessage(threadId, messageId)
	if err != nil {
		fmt.Println("MESSAGER ADAPTER GET MESSAGE ERROR: ", err)
		return err
	}

	comp, err := a.AIAdapter.CreateCompletion(m.Text)
	if err != nil {
		fmt.Println("AI ADAPTER CREATE COMPLETION ERROR: ", err)
		return err
	}

	err = a.MessagerAdapter.SendMessage(threadId, "MESSAGE", comp.Choices[0].Message.Content, m.ChannelID, m.ChannelAccountID, "A-64307862")

	return err
}
