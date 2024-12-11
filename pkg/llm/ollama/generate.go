package ollama

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/fenek-dev/llm-craft/internal/entity"
	"github.com/ollama/ollama/api"
)

func (o *Ollama) Generate(ctx context.Context, el1, el2 string) (entity.Element, error) {

	b, err := json.Marshal(entity.Element{})
	if err != nil {
		return entity.Element{}, err
	}

	prompt := fmt.Sprintf(`
	You are an intelligent assistant that helps combine elements in a crafting system.
	Combinantion result should be logical and make sense.
	
	Here is how combinations work:
	- Two elements can combine to form a new element. For example, "fire" + "water" = "steam".
	- Always response with an emoji that represents the result.
	
	Your task:
	1. Take two elements and combine them into a new element.
	3. Return the response in JSON format, with the following structure:
	   {
		 "name": "<name>",
		 "emoji": "<emoji>",
	   }
	
	Now, combine the following elements:
	- Element 1: %s
	- Element 2: %s
	
	Provide your response in the requested JSON format.
	YOU MUST ALWAYS GENERATE EMOJI.
	`, el1, el2)

	req := &api.GenerateRequest{
		Model:  o.model,
		Prompt: prompt,
		Format: b,
		Stream: new(bool),
		Options: map[string]interface{}{
			"temperature": 0.2,
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
	err = o.client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}

	r := <-c
	fmt.Println("r", r)

	return r, nil
}
