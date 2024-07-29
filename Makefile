.PHONY: build start-app run clean

# Variables
BINARY_NAME := firstpass
BUILD_DIR := .bin
MAIN_FILE := main.go

build:
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

start-app: build
	@$(BUILD_DIR)/$(BINARY_NAME) $(ARGS)

run: build
	@$(BUILD_DIR)/$(BINARY_NAME) $(ARGS)

clean:
	@rm -rf $(BUILD_DIR)/$(BINARY_NAME)
