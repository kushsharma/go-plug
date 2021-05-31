GOVERSION := $(shell go version | cut -d ' ' -f 3 | cut -d '.' -f 2)

.PHONY: build build-plugin check fmt lint test test-race vet test-cover-html help generate-proto install
.DEFAULT_GOAL := help

build: generate-proto build-plugin ## build all
	@echo " > building core"
	go build -o goplug .
	@echo " > build finished"

build-plugin: ## build plugins
	@echo " > building plugins"
	go build -o ./plugin-sql ./plugins/sql/sql.go

generate-proto: ## Generate proto files
	@echo " > generating protos"
	@buf generate

check: test-race fmt vet lint ## Run tests and linters

test: ## Run tests
	go test ./... -race

test-race: ## Run tests with race detector
	go test -race ./...

clean :
	rm -rf dist

fmt: ## Run gofmt linter
ifeq "$(GOVERSION)" "12"
	@for d in `go list` ; do \
		if [ "`gofmt -l -s $$GOPATH/src/$$d | tee /dev/stderr`" ]; then \
			echo "^ improperly formatted go files" && echo && exit 1; \
		fi \
	done
endif

lint: ## Run golint linter
	@for d in `go list` ; do \
		if [ "`golint $$d | tee /dev/stderr`" ]; then \
			echo "^ golint errors!" && echo && exit 1; \
		fi \
	done

vet: ## Run go vet linter
	@if [ "`go vet | tee /dev/stderr`" ]; then \
		echo "^ go vet errors!" && echo && exit 1; \
	fi

test-cover-html: ## Generate test coverage report
	go test -coverprofile=coverage.out -covermode=count
	go tool cover -func=coverage.out

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Install required libraries
	@echo "> installing dependencies"
	go get -u github.com/golang/protobuf/proto@v1.4.3
	go get -u github.com/golang/protobuf/protoc-gen-go@v1.4.3
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go get -u google.golang.org/grpc@v1.35.0
	go get -u github.com/bufbuild/buf/cmd/buf@v0.37.0