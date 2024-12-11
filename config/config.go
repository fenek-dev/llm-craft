package config

import "github.com/fenek-dev/sdk/pkg/db/redis"

type Config struct {
	Redis redis.Config `yaml:"redis" env-prefix:"REDIS_"`
}
