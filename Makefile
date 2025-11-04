.PHONY: run build clean tidy debug compose-up compose-down compose-logs compose-rebuild 

run:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go run -v -race ./cmd/server/main.go

tidy:
	go mod tidy && go mod vendor

clean:
	rm -rf bin/ coverage.out coverage.html

debug:
	go run -race ./cmd/server/main.go

compose-up:
	docker compose up -d

compose-down:
	docker compose down

compose-logs:
	docker compose logs -f

compose-rebuild:
	docker compose up -d --build

security-scan:
	@echo "Scanning filesystem"
	trivy fs .

security-docker:
	@echo "Scanning Docker image"
	docker compose build api
	trivy image mercedes-api

security-secrets: 
	@echo "scanning secrets"
	trivy fs --scanners secret .

