.PHONY: sandbox
sandbox:
	go run ./cmd/sandbox/main.go

.PHONY: test
test:
	go test -cover ./...