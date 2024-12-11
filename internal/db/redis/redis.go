package redis

import (
	"context"
	"log/slog"
	"time"

	"github.com/fenek-dev/llm-craft/config"
	rdb "github.com/fenek-dev/sdk/pkg/db/redis"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	conn *redis.Client

	log *slog.Logger
}

func New(ctx context.Context, cfg *config.Config, opts ...Option) *Redis {
	conn := rdb.MustNew(cfg.Redis)

	rdb.MustPing(ctx, conn, 5*time.Second)

	rd := &Redis{conn: conn}

	for _, opt := range opts {
		opt(rd)
	}

	return rd
}

func (r *Redis) Close() error {
	return r.conn.Close()
}
