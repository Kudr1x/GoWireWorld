lint:
	golangci-lint run ./...
	@echo "ok"

test:
	go test ./src/game/...
	@echo "ok"