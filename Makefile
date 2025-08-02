BINARY_NAME=server
BUILD_DIR=build

.PHONY: all build test clean

all: build

build-server:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./server

test:
	go test ./...

clean:
	rm -rf $(BUILD_DIR)