package redis

import "log/slog"

type Option func(*Redis)

func WithLogger(logger *slog.Logger) Option {
	return func(r *Redis) {
		r.log = logger
	}
}
