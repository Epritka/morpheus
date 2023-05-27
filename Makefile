.PHONY: sandbox
sandbox:
	go run ./sandbox/main.go

.PHONY: test
test:
	go test -cover ./...