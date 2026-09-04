BINARY := go-test-skeleton
PORT ?= 3000

.PHONY: all build test cover run fmt vet check tidy clean

all: check build

build:
	go build -o bin/$(BINARY) .

test:
	go test ./... -race -cover

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

run:
	go run . serve --port $(PORT)

fmt:
	gofmt -l -w .

vet:
	go vet ./...

# check is what CI runs, and what you should run before submitting.
check:
	test -z "$$(gofmt -l .)" || { echo "gofmt: files need formatting:"; gofmt -l .; exit 1; }
	go vet ./...
	go test ./... -race

tidy:
	go mod tidy

clean:
	rm -rf bin coverage.out
