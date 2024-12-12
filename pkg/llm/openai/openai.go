package openai

import (
	"context"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	client *openai.Client
	model  string
	prompt string
}

func New(apiKey string, opts ...Option) *OpenAI {
	client := openai.NewClient(apiKey)
	o := &OpenAI{
		client: client,
		model:  openai.GPT4oMini,
		prompt: prompt,
	}

	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (o *OpenAI) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: o.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: "ping",
				},
			},
		},
	)
	return err
}
