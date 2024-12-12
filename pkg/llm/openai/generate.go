package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/fenek-dev/llm-craft/internal/entity"
	openai "github.com/sashabaranov/go-openai"
)

func (o *OpenAI) Generate(ctx context.Context, el1, el2 string) (entity.Element, error) {
	resp, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: o.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: o.prompt,
				},
				{
					Role:    "user",
					Content: fmt.Sprintf("%s + %s = ?", el1, el2),
				},
			},
			Temperature: 0,
		},
	)
	if err != nil {
		return entity.Element{}, err
	}

	result := entity.Element{}
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result); err != nil {
		return entity.Element{}, err
	}

	log.Println(el1, "+", el2, "=", result.Name, result.Emoji)
	return result, nil
}
