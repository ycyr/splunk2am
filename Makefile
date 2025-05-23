BINARY_NAME=splunk2am
BUILD_DIR=cmd/server
VERSION=$(shell date +'%y%m%d-%H%M')

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X github.com/ycyr/splunk2alertmanager/pkg/config.version=$(VERSION)" -o splunk2am ./cmd/server

clean:
	rm -f $(BINARY_NAME)

.PHONY: build clean

