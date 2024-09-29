# Makefile for cross-compiling webfoot for multiple platforms

BINARY_NAME=webfoot
BUILD_DIR=bin

.PHONY: all clean build-linux build-mac build-windows build-all

# Default target: build for all platforms
all: build-all

# Build for Linux
build-linux:
	@mkdir -p $(BUILD_DIR)/linux
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/linux/$(BINARY_NAME) ./cmd/main.go
	@echo "Built Linux binary in $(BUILD_DIR)/linux/$(BINARY_NAME)"

# Build for macOS
build-mac:
	@mkdir -p $(BUILD_DIR)/mac
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/mac/$(BINARY_NAME) ./cmd/main.go
	@echo "Built macOS binary in $(BUILD_DIR)/mac/$(BINARY_NAME)"

# Build for Windows
build-windows:
	@mkdir -p $(BUILD_DIR)/windows
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/windows/$(BINARY_NAME).exe ./cmd/main.go
	@echo "Built Windows binary in $(BUILD_DIR)/windows/$(BINARY_NAME).exe"

# Build all binaries for Linux, macOS, and Windows
build-all: build-linux build-mac build-windows
	@echo "Built all binaries."

# Clean up the build artifacts
clean:
	rm -rf $(BUILD_DIR)
	@echo "Cleaned up build artifacts."
