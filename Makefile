BINARY_NAME=splunk2am
BUILD_DIR=cmd/server
VERSION=$(shell date +'%y%m%d-%H%M')

build:
	go build -ldflags "-X main.version=$(VERSION)" -o $(BINARY_NAME) ./$(BUILD_DIR)

clean:
	rm -f $(BINARY_NAME)

.PHONY: build clean

