all: install

dependencies:
	go get -u github.com/golang/lint/golint
	go get -u golang.org/x/tools/cmd/cover
	go get -u golang.org/x/tools/cmd/vet

fmt:
	go fmt ./...

install: fmt
	go install ./...

lint:
	golint ./...

vet:
	go vet ./...

test: fmt lint vet
	go test ./...

cover: fmt lint vet
	@go test -covermode=atomic -coverprofile=cover.out . \
	&& go tool cover -html=cover.out -o=cover.html \
	&& rm cover.out
cover-integration: fmt lint vet
	@go test -run=TestIntegration -covermode=atomic -coverprofile=cover.out . \
	&& go tool cover -html=cover.out -o=cover.html \
	&& rm cover.out
cover-unit: fmt lint vet
	@go test -run=TestUnit -covermode=atomic -coverprofile=cover.out . \
	&& go tool cover -html=cover.out -o=cover.html \
	&& rm cover.out

.PHONY: all dependencies fmt install test cover vet
