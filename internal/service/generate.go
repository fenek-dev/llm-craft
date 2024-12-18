package service

import (
	"context"

	"github.com/fenek-dev/llm-craft/internal/entity"
)

func (s *Service) Generate(ctx context.Context, el1, el2 string) (entity.Element, bool, error) {
	//result, err := s.cache.Get(ctx, el1, el2)
	//if err == nil {
	//	return result, false, nil
	//}

	result, err := s.llm.Generate(ctx, el1, el2)

	if err != nil {
		return entity.Element{}, false, err
	}

	if result.Name == "" {
		return entity.Element{}, false, nil
	}
	//
	//err = s.cache.Set(ctx, el1, el2, result)
	//if err != nil {
	//	s.log.Error("failed to cache result", logger.Err(err))
	//}

	return result, true, nil
}
