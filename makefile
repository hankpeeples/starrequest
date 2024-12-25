.PHONY: run test build clean

# Default target
all: run

# Run the application
run:
	@go mod tidy
	@go run starrequest

# Run tests
test:
	go test -cover ./...

# Clean build artifacts
clean:
	go clean
	rm -f starrequest