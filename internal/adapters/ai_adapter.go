package adapters

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type AIAdapter struct {
	Client    *openai.Client
	Ctx       context.Context
	Model     string
	MaxTokens int
}

func NewAIAdapter(aiAPIKey, model string, maxTokens int) *AIAdapter {
	ctx := context.Background()
	client := openai.NewClient(aiAPIKey)

	return &AIAdapter{
		Client:    client,
		Ctx:       ctx,
		Model:     model,
		MaxTokens: maxTokens,
	}
}

func (a *AIAdapter) CreateCompletion(prompt string) (*openai.ChatCompletionResponse, error) {
	req := openai.ChatCompletionRequest{
		Model:     a.Model,
		MaxTokens: a.MaxTokens,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	comp, err := a.Client.CreateChatCompletion(a.Ctx, req)
	return &comp, err
}
