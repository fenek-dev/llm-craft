package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/fenek-dev/llm-craft/internal/entity"
)

func (r *Redis) Get(ctx context.Context, el1, el2 string) (entity.Element, error) {

	els := []string{el1, el2}

	sort.Strings(els)

	key := fmt.Sprintf("combine:%s:%s", els[0], els[1])

	cmd := r.conn.Get(ctx, key)

	if cmd.Err() != nil {
		return entity.Element{}, cmd.Err()
	}

	var el entity.Element

	err := json.Unmarshal([]byte(cmd.Val()), &el)
	if err != nil {
		return entity.Element{}, err
	}

	return el, nil
}

func (r *Redis) Set(ctx context.Context, el1, el2 string, result entity.Element) error {

	els := []string{el1, el2}

	sort.Strings(els)

	key := fmt.Sprintf("combine:%s:%s", els[0], els[1])

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	cmd := r.conn.Set(ctx, key, resultJSON, 0)

	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
