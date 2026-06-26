fmt:
	golangci-lint fmt

lint:
	golangci-lint run

.PHONY: fmt, lint
