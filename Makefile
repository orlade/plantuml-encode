.PHONY: all tidy lint build test releasebuild docker docker.run

VERSION=`cat VERSION`
GO_VERSION=$(shell go version | cut -d' ' -f3)

LDFLAGS='-X "main.Version=$(VERSION)" -X "main.GoVersion=$(GO_VERSION)"'

all: tidy build

.PHONY: tidy
tidy:
	go mod tidy
	goimports -w .
	gofmt -w .

lint:
	golangci-lint run ./...

build: main.go
	go build -v -o ./dist/plantuml-encode -ldflags=$(LDFLAGS) main.go

test: main.go main_test.go
	go test -v -race ./...

releasebuild: main.go
	GOVERSION=$(GOVERSION) goreleaser build --rm-dist --snapshot

docker: main.go
	docker build -t orlade/plantuml-encode .

# Pipe input to this target to encode it.
docker.run:
	docker run -i --rm orlade/plantuml-encode
