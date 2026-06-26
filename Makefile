fmt:
	golangci-lint fmt

lint:
	golangci-lint run

test-all:
	go test -v ./...

.PHONY: fmt, lint
