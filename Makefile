.PHONY: run build clean tidy

run:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go run -v -race ./cmd/server/main.go

tidy:
	go mod tidy && go mod vendor

clean:
	rm -rf bin/ coverage.out coverage.html