package service

import (
	"context"
	"log/slog"

	"github.com/fenek-dev/llm-craft/internal/entity"
)

type Service struct {
	cache Cache
	llm   LLM
	log   *slog.Logger
}

type Cache interface {
	Get(ctx context.Context, el1, el2 string) (entity.Element, error)
	Set(ctx context.Context, el1, el2 string, result entity.Element) error
}

type LLM interface {
	Generate(ctx context.Context, el1, el2 string) (entity.Element, error)
}

func New(cache Cache, llm LLM) *Service {
	return &Service{
		cache: cache,
		llm:   llm,
	}
}
