VERSION=`cat VERSION`
GO_VERSION=$(strip $(subst go version, , $(shell go version)))

LDFLAGS='-X "main.Version=$(VERSION)" -X "main.GoVersion=$(GO_VERSION)"'

all: tidy build

.PHONY: tidy
tidy:
	go mod tidy
	goimports -w .
	gofmt -w .

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: build
build: main.go
	go build -v -o ./dist/plantuml-encode -ldflags=$(LDFLAGS) main.go

.PHONY: test
test: main.go main_test.go
	go test ./...

.PHONY: docker
docker: main.go
	docker build -t orlade/plantuml-encode .

# Pipe input to this target to encode it.
.PHONY: docker.run
docker.run:
	docker run -i --rm orlade/plantuml-encode
