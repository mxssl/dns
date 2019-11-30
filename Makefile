BINARY_NAME=dns
CURRENT_DIR=$(shell pwd)
TAG=$(shell git name-rev --tags --name-only $(shell git rev-parse HEAD))
DOCKER_REGISTRY=mxssl
export GO111MODULE=on

.PHONY: all build clean test build-linux build-darwin build-windows github-release-dry github-release docker-release

all: test build

build:
	go build -mod=vendor -o ${BINARY_NAME} -v

clean:
	rm -f ${BINARY_NAME}

lint:
	golangci-lint run -v

# Cross compilation
build-linux:
	env GOOS=linux \
	GOARCH=amd64 \
	go build \
	-o ${BINARY_NAME}-linux-amd64
build-darwin:
	env GOOS=darwin \
	GOARCH=amd64 \
	go build \
	-o ${BINARY_NAME}-darwin-amd64
build-windows:
	env GOOS=windows \
	GOARCH=amd64 \
	go build \
	-o ${BINARY_NAME}-windows-amd64

github-release-dry:
	@echo "TAG: ${TAG}"
	goreleaser release --rm-dist --snapshot --skip-publish

github-release:
	@echo "TAG: ${TAG}"
	goreleaser release --rm-dist

docker-release:
	@echo "Registry: ${DOCKER_REGISTRY}"
	@echo "TAG: ${TAG}"
	docker build --tag ${DOCKER_REGISTRY}/${BINARY_NAME}:${TAG} .
	docker push ${DOCKER_REGISTRY}/${BINARY_NAME}:${TAG}
