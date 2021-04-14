ARCH ?= amd64
OS ?= linux
BUILD_IMAGE ?= golang:1.16
GOPATH ?= $(shell go env GOPATH)
GOPATH_SRC := $(GOPATH)/src/
CURRENT_WORK_DIR := $(shell pwd)
ALL_FILES=$(shell find . -path ./vendor -prune -type f -o -name *.proto)
PKG=github.com/leachim2k/go-shorten

GIT_COMMIT := $(shell git rev-parse HEAD)
VERSION ?= $(GIT_COMMIT)

TAG ?= $(VERSION)
IMAGE ?= ""
YES ?= ""

all: build

clean: clean-dirs

container: check-image build
	@docker build --quiet -t $(IMAGE):$(TAG) -f hack/release/Dockerfile .
	$(info container image built $(IMAGE):$(TAG))

ifeq ($(YES), 1)
push-container: check-image container
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker push $(IMAGE):$(TAG)
else
push-container:
	$(warning push disabled. to enable set environment YES=1)
endif

ifndef IMAGE
check-image:
	  $(error env IMAGE is undefined)
else
check-image:
	  $(info target image is $(IMAGE))
endif

build: build-docs build-webui $(subst cmd, dist/$(ARCH), $(wildcard cmd/*))

dist/$(ARCH)/%: build-dirs
	$(info building binary $(notdir $@))
	@docker run \
		--rm \
		-u $$(id -u):$$(id -g) \
		-v "$$(pwd):/src" \
		-v "$$(pwd)/dist/$(OS)/$(ARCH):/go/bin" \
		-v "$$(pwd)/.gocache/:/go/cache" \
		-w /src \
		$(BUILD_IMAGE) \
		/bin/sh -c " \
			ARCH=$(ARCH) \
			OS=$(OS) \
			VERSION=$(VERSION) \
			COMMIT=$(GIT_COMMIT) \
			PKG=$(PKG) \
			BIN=$(notdir $@) \
			GO111MODULE=on \
			./hack/build.sh \
		"

test: build-dirs
	$(info run test)
	@docker run \
		--rm \
		-u $$(id -u):$$(id -g) \
		-v "$$(pwd):/src" \
		-v "$$(pwd)/dist/$(OS)/$(ARCH):/go/bin" \
		-v "$$(pwd)/.gocache/:/go/cache" \
		-w /src \
		$(BUILD_IMAGE) \
		/bin/sh -c "CGO_ENABLED=0 GO111MODULE=on GOCACHE=/go/cache go test -mod=vendor -v ./..."

build-dirs:
	@echo "build-dirs"
	@mkdir -p ./dist/$(OS)/$(ARCH)
	@mkdir -p ./bin
	@mkdir -p ./.gocache

clean-dirs:
	$(info clean up cache and dist folders)
	@rm -rf ./bin
	@rm -rf ./docs
	@rm -rf ./dist
	@rm -rf ./webui/build
	@rm -rf ./.gocache

build-webui:
	@echo "build-webui"
	@npm install --prefix webui
	@npm run build --prefix webui
ifeq (,$(wildcard ./bin/go-bindata))
	@echo "Download and build go-bindata"
	@env GO111MODULE=off go get github.com/go-bindata/go-bindata/go-bindata
	@env GO111MODULE=off go build -o ./bin/go-bindata github.com/go-bindata/go-bindata/go-bindata
endif
	@./bin/go-bindata -fs -o pkg/server/webui.data.go -prefix "webui/build" -pkg server webui/build/...

build-docs:
	@echo "build-docs"
ifeq (,$(wildcard ./bin/swag))
	@echo "Download and build swag"
	@env GO111MODULE=off go get github.com/swaggo/swag/cmd/swag
	@env GO111MODULE=off go build -o ./bin/swag github.com/swaggo/swag/cmd/swag
endif
	@./bin/swag init -g cmd/shorten/main.go
