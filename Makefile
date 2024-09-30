# Variables
APP_NAME = task-manager
CMD_DIR = ./cmd

# Default target
.PHONY: all
all: build

# Build the project
.PHONY: build
build:
	go build -o $(APP_NAME) $(CMD_DIR)

# Run the project
.PHONY: run
run:
	go run $(CMD_DIR)/main.go

# Clean the build
.PHONY: clean
clean:
	rm -f $(APP_NAME)
