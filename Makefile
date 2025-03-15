.PHONY: build run test clean

# Default build target
build:
	go build -o honeygo ./cmd/honeygo

# Run the honeypot
run: build
	./honeygo

# Run with custom config
run-config: build
	./honeygo -config $(CONFIG)

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Generate SSH key (RSA 2048 bits)
generate-key:
	mkdir -p keys
	ssh-keygen -t rsa -b 2048 -f keys/honeygo_key -N ""

# Clean build artifacts
clean:
	rm -f honeygo
	rm -f coverage.out coverage.html

# Set up initial config
init:
	mkdir -p logs keys
	[ -f config.yaml ] || cp config.example.yaml config.yaml

# Help target
help:
	@echo "HoneyGO Makefile targets:"
	@echo "  build         - Build the honeypot binary"
	@echo "  run           - Build and run with default config"
	@echo "  run-config    - Build and run with specified config (CONFIG=path)"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  generate-key  - Generate SSH host key"
	@echo "  clean         - Remove build artifacts"
	@echo "  init          - Set up initial directories and config"