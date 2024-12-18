
# Makefile for Dynamic Notification System

# Variables
GO := go
BUILD_DIR := build
MAIN_FILE := main.go
OUTPUT_BINARY := notification-system
PLUGIN_DIR := plugins

# Targets
.PHONY: all clean build build-plugins

# Default target
all: build build-plugins

# Initialize the Go packages
init:
	go mod init dynamic-notification-system
	go mod tidy

# Build the main binary
build:
	@echo "Building main application..."
	$(GO) build -o $(BUILD_DIR)/$(OUTPUT_BINARY) $(MAIN_FILE)

# Build plugins
build-plugins: 
	go build -buildmode=plugin -o build/plugins/telegram.so plugins/telegram/telegram.go
	go build -buildmode=plugin -o build/plugins/discord.so plugins/discord/discord.go
	go build -buildmode=plugin -o build/plugins/slack.so plugins/slack/slack.go
	go build -buildmode=plugin -o build/plugins/teams.so plugins/teams/teams.go
	go build -buildmode=plugin -o build/plugins/webhook.so plugins/webhook/webhook.go
	go build -buildmode=plugin -o build/plugins/smtp.so plugins/smtp/smtp.go
	go build -buildmode=plugin -o build/plugins/push.so plugins/push/push.go
	go build -buildmode=plugin -o build/plugins/sms.so plugins/sms/sms.go
	go build -buildmode=plugin -o build/plugins/signal.so plugins/signal/signal.go
	go build -buildmode=plugin -o build/plugins/rocketchat.so plugins/rocketchat/rocketchat.go
	go build -buildmode=plugin -o build/plugins/nfyt.so plugins/nfyt/nfyt.go


# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)

# Help command
help:
	@echo "Makefile for Dynamic Notification System"
	@echo ""
	@echo "Targets:"
	@echo "  all             - Build the main application and plugins"
	@echo "  init            - Initialize the Go packages"
	@echo "  build           - Build the main application"
	@echo "  build-plugins   - Build all plugins"
	@echo "  clean           - Clean all build artifacts"
