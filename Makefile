GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
LDFLAGS = -X "main.buildCommit=$(COMMIT)"

BINARY_NAME=drone-newrelic

all: build image push

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o build/$(BINARY_NAME) main.go plugin.go

.PHONY: clean
clean:
	rm -rf build

.PHONY: image
image:
	@docker build -t kuperiu/drone-newrelic:latest .

.PHONY: push
push:
	@docker push kuperiu/drone-newrelic:latest