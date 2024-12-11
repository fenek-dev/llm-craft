package main

import (
	"context"
	"net/http"
	url2 "net/url"

	"github.com/fenek-dev/llm-craft/config"
	"github.com/fenek-dev/llm-craft/internal/db/redis"
	httpServer "github.com/fenek-dev/llm-craft/internal/http"
	"github.com/fenek-dev/llm-craft/internal/service"
	"github.com/fenek-dev/llm-craft/pkg/llm/ollama"
	sdkconfig "github.com/fenek-dev/sdk/pkg/config"
)

func main() {
	var (
		ctx = context.Background()
		cfg = &config.Config{}
	)

	sdkconfig.MustParse(cfg, ".env")

	rdb := redis.New(ctx, cfg)

	url, err := url2.Parse("http://localhost:11434")
	if err != nil {
		panic(err)
	}

	llm := ollama.New(url, http.DefaultClient)

	if err = llm.Ping(ctx); err != nil {
		panic(err)
	}

	srv := service.New(rdb, llm)

	server := httpServer.New(srv)

	server.InitRoutes()

	server.Run(ctx, ":8080")
}
