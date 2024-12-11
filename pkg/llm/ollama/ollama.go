package ollama

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/ollama/ollama/api"
)

type Ollama struct {
	client *api.Client
	model  string
}

func New(base *url.URL, http *http.Client, opts ...Option) *Ollama {
	client := api.NewClient(base, http)

	ollama := &Ollama{
		client: client,
		model:  "llama3.2",
	}

	for _, opt := range opts {
		opt(ollama)
	}

	return ollama
}

func (o *Ollama) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := o.client.Version(ctx)
	return err
}
