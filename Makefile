.PHONY: run build clean tidy debug test test-coverage coverage-html swagger compose-up compose-down compose-logs compose-rebuild security-scan security-docker security-secrets

swagger:
	@echo "Generating Swagger documentation..."
	@swag init -g cmd/server/main.go --output ./docs --parseDependency --parseInternal
	@echo "✅ Swagger docs generated! View at: http://localhost:6969/swagger/index.html"

run:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go run -v -race ./cmd/server/main.go

build:
	CGO_ENABLED=0 go build -o bin/server ./cmd/server/main.go

tidy:
	go mod tidy && go mod vendor

clean:
	rm -rf bin/ coverage.out coverage.html

test:
	go test -v -race ./...

test-coverage:
	go test -v -race -coverprofile=coverage.out -covermode=atomic \
		$(shell go list ./... | grep -v -e '/cmd/' -e '/mocks' -e '/dto.go')
	go tool cover -func=coverage.out | grep -v -e 'main.go' -e 'mock.go'

coverage-html:
	go test -v -race -coverprofile=coverage.out -covermode=atomic \
		$(shell go list ./... | grep -v -e '/cmd/' -e '/mocks')
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html (excludes: cmd, mocks)"

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

