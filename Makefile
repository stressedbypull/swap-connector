.PHONY: run build clean tidy debug docker-up docker-down docker-logs docker-rebuild docker-health

run:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go run -v -race ./cmd/server/main.go

tidy:
	go mod tidy && go mod vendor

clean:
	rm -rf bin/ coverage.out coverage.html

debug:
	go run -race ./cmd/server/main.go

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f

docker-rebuild:
	docker compose up -d --build

docker-health:
	@echo "=== Container Status ==="
	@docker compose ps
	@echo "\n=== Database Health Check ==="
	@docker compose exec db pg_isready -U postgres -d swap_connector || echo "DB not ready"
	@echo "\n=== API Health Check ==="
	@curl -s http://localhost:6969/ping || echo "API not ready"