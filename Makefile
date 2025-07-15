APP_NAME := divo
BUILD_DIR := bin
PORT ?= 3000

# Detect OS and set binary name
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	BIN := $(BUILD_DIR)/$(APP_NAME)
endif
ifeq ($(UNAME_S),Darwin)
	BIN := $(BUILD_DIR)/$(APP_NAME)-mac
endif
ifeq ($(OS),Windows_NT)
	BIN := $(BUILD_DIR)/$(APP_NAME).exe
endif

.PHONY: all build run clean cross

all: build

## Build for current OS
build:
	@echo "[?] Detected OS: $(UNAME_S)"
	@echo "[&] Building $(APP_NAME) for current platform..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BIN) ./cmd/$(APP_NAME)
	@echo "[+] Built binary: $(BIN)"

## Run the proxy with default or given port
run: build
	@echo "[^] Running: $(BIN) http $(PORT)"
	@$(BIN) http $(PORT)

## Clean generated files
ifeq ($(OS),Windows_NT)
clean:
	@echo "[-] Cleaning..."
	@powershell -Command "Remove-Item -Recurse -Force $(BUILD_DIR)"
else
clean:
	@echo "[-] Cleaning..."
	@rm -rf $(BUILD_DIR)
endif

## Cross-compile for Linux, Windows, and macOS
cross:
	@mkdir -p $(BUILD_DIR)
	@echo "[*]  Cross-compiling..."
	GOOS=linux GOARCH=amd64   go build -o $(BUILD_DIR)/$(APP_NAME)        ./cmd/$(APP_NAME)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME).exe    ./cmd/$(APP_NAME)
	GOOS=darwin GOARCH=amd64  go build -o $(BUILD_DIR)/$(APP_NAME)-mac    ./cmd/$(APP_NAME)
	@echo "[+] Binaries built in ./$(BUILD_DIR):"
	@ls -lh $(BUILD_DIR)
