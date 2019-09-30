GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
BUILD_DIR ?= ./build

ORG := github.com/platform9
REPOPATH ?= $(ORG)/alertmgr-config

DOCKER_IMAGE_NAME = platform9/alertmgrcfg-operator
DOCKER_IMAGE_TAG ?= latest

LDFLAGS := -s -w -extldflags '-static'

SRCFILES := $(shell find ./pkg)

test:
	go test ./pkg/...

build/bin/alertmgr-config-operator:  build/bin/alertmgr-config-operator-$(GOOS)-$(GOARCH)
	cp build/bin/alertmgr-config-operator-$(GOOS)-$(GOARCH) build/bin/alertmgr-config-operator

build/bin/alertmgr-config-operator-darwin-amd64: $(SRCFILES)
	GOARCH=amd64 GOOS=darwin go build --installsuffix cgo -a -o build/bin/alertmgr-config-operator-darwin-amd64 cmd/manager/main.go

build/bin/alertmgr-config-operator-linux-amd64: $(SRCFILES)
	GOARCH=amd64 GOOS=linux go build --installsuffix cgo -a -o build/bin/alertmgr-config-operator-linux-amd64 cmd/manager/main.go


.PHONY: clean
clean:
	rm -fr build/

.PHONY: binary
binary: build/bin/alertmgr-config-operator

.PHONY: image
image: test build/bin/alertmgr-config-operator-linux-amd64
	docker build -t $(DOCKER_IMAGE_NAME) .
