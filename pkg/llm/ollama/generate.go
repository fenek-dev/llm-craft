package ollama

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fenek-dev/llm-craft/internal/entity"
	"github.com/ollama/ollama/api"
)

func (o *Ollama) Generate(ctx context.Context, el1, el2 string) (entity.Element, error) {

	req := &api.GenerateRequest{
		Model:  o.model,
		Prompt: fmt.Sprintf("%s + %s = ?", el1, el2),
		System: o.prompt,
		Stream: new(bool),
		Format: o.format,
		KeepAlive: &api.Duration{
			Duration: time.Minute,
		},
		Options: map[string]interface{}{
			"temperature": 0,
		},
	}

	c := make(chan entity.Element, 1)

	respFunc := func(resp api.GenerateResponse) error {
		// Only print the response here; GenerateResponse has a number of other
		// interesting fields you want to examine.
		defer close(c)
		result := entity.Element{}
		if err := json.Unmarshal([]byte(resp.Response), &result); err != nil {
			return err
		}
		c <- result
		return nil
	}
	err := o.client.Generate(ctx, req, respFunc)
	if err != nil {
		return entity.Element{}, err
	}

	select {
	case <-ctx.Done():
		return entity.Element{}, ctx.Err()
	case r := <-c:
		log.Println(el1, "+", el2, "=", r.Name, r.Emoji)
		return r, nil
	}

}
