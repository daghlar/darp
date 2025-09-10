# DARP Makefile
# Cloudflare WARP client for Arch Linux

.PHONY: build clean test install uninstall help

# Variables
APP_NAME = darp
VERSION = $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME = $(shell date -u '+%Y-%m-%d_%H:%M:%S')
BUILD_DIR = build
BINARY_NAME = $(APP_NAME)-$(VERSION)

# Go build flags
LDFLAGS = -ldflags "-X main.version=$(VERSION) -X main.build=$(BUILD_TIME)"

# Default target
all: build

# Build the application
build:
	@echo "🔨 Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME) ./cmd/darp
	@echo "✅ Build completed: $(BUILD_DIR)/$(APP_NAME)"

# Build for multiple architectures
build-all:
	@echo "🔨 Building for multiple architectures..."
	@mkdir -p $(BUILD_DIR)
	@echo "  Building for linux/amd64..."
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/darp
	@echo "  Building for linux/arm64..."
	@GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 ./cmd/darp
	@echo "✅ Multi-arch build completed"

# Run tests
test:
	@echo "🧪 Running tests..."
	@go test ./... -v

# Run tests with coverage
test-coverage:
	@echo "🧪 Running tests with coverage..."
	@go test ./... -cover

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "✅ Clean completed"

# Install dependencies
deps:
	@echo "📦 Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo "✅ Dependencies installed"

# Format code
fmt:
	@echo "🎨 Formatting code..."
	@go fmt ./...
	@echo "✅ Code formatted"

# Lint code
lint:
	@echo "🔍 Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not found, skipping linting"; \
	fi

# Run the application
run: build
	@echo "🚀 Running $(APP_NAME)..."
	@sudo $(BUILD_DIR)/$(APP_NAME) $(ARGS)

# Create distribution package
package: build-all
	@echo "📦 Creating distribution package..."
	@./build.sh
	@echo "✅ Package created: $(BUILD_DIR)/$(BINARY_NAME).tar.gz"

# Install to system (requires root)
install: build
	@echo "📦 Installing $(APP_NAME) to system..."
	@sudo cp $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/
	@sudo chmod +x /usr/local/bin/$(APP_NAME)
	@echo "✅ $(APP_NAME) installed to /usr/local/bin/"

# Uninstall from system
uninstall:
	@echo "🗑️  Uninstalling $(APP_NAME) from system..."
	@sudo rm -f /usr/local/bin/$(APP_NAME)
	@echo "✅ $(APP_NAME) uninstalled"

# Show help
help:
	@echo "DARP - Cloudflare WARP Client for Arch Linux"
	@echo "=============================================="
	@echo ""
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  build-all    - Build for multiple architectures"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Install dependencies"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code"
	@echo "  run          - Build and run the application"
	@echo "  package      - Create distribution package"
	@echo "  install      - Install to system (requires root)"
	@echo "  uninstall    - Uninstall from system"
	@echo "  help         - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make build                    # Build the application"
	@echo "  make run ARGS=\"status\"        # Run with arguments"
	@echo "  make package                  # Create distribution package"
	@echo "  sudo make install             # Install to system"
