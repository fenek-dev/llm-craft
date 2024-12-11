run:
	go run cmd/main.go

app-up:
	docker compose -f=docker/docker-compose.yml -p=llm-craft up

redis-up:
	docker compose -f=docker/docker-compose.redis.yml -p=redis-llm-craft up