# Binary name
BINARY=search

# Go commands
GO=go
BUILD=$(GO) build
RUN=$(GO) run
TEST=$(GO) test

# Default target
.PHONY: all
all: build

.PHONY: build
build:
	go build -o ./bin/dupscanner ./cmd

.PHONY: run
run:
	go run ./cmd --dir ./ 

.PHONY: run-custom
run-custom:
	go run ./cmd $(ARGS)

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY)

# Format code
.PHONY: fmt
fmt:
	$(GO) fmt ./...

# Lint (requires golangci-lint installed)
.PHONY: lint
lint:
	golangci-lint run

# Test
.PHONY: test
test:
	$(TEST) ./...

# Race detection (important for concurrency)
.PHONY: race
race:
	$(TEST) -race ./...

# Benchmark (if you add benchmarks later)
.PHONY: bench
bench:
	$(TEST) -bench=. ./...

# Install binary to GOPATH/bin
.PHONY: install
install:
	$(GO) install .

# Help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build        - Build the binary"
	@echo "  make run          - Run with default args"
	@echo "  make run-custom   - Run with custom args (use ARGS='...')"
	@echo "  make clean        - Remove binary"
	@echo "  make fmt          - Format code"
	@echo "  make lint         - Run linter"
	@echo "  make test         - Run tests"
	@echo "  make race         - Run tests with race detector"
	@echo "  make install      - Install binary"