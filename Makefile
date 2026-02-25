lint:
	golangci-lint run ./...
	@echo "ok"

test:
	go test ./src/core/...
	@echo "ok"