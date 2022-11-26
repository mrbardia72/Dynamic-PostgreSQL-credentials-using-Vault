LOGFILE=$(LOGPATH) `date +'%A-%b-%d-%Y'`
branch := $(shell git branch --show-current)
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export GO111MODULE=on

.PHONY: help
help: ## Shows help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.PHONY: build-app
build-app:
	CGO_ENABLED=1 go build -v -o $(ROOT)/bin/main -ldflags="-s -w" $(ROOT)/cmd/api/*.go

.which-go:
	@which go > /dev/null || (echo "install go from https://golang.org/dl/" & exit 1)

format: .which-go
	gofmt -s -w $(ROOT)

.which-lint:
	@which golangci-lint > /dev/null || (echo "install golangci-lint from https://github.com/golangci/golangci-lint" & exit 1)

lint: .which-lint
	golangci-lint run

clean: # run make format and make lint
	gofmt -s -w $(ROOT)
	golangci-lint run

.PHONY: test
test: .which-go ## Tests go files
	CGO_ENABLED=1 go test -race -coverprofile=coverage.txt -covermode=atomic $(ROOT)/... -v

.PHONY: cm
cm: ## 🌱 git commit
	git add .
	git commit -m "$(branch)-${LOGFILE}"
	git push origin $(branch)

.PHONY: test-flag
test-flag: ## test with flags
	go test -tags=author_logic_test -bench=. -benchmem -v -cover ./...
	go test -tags=sql_mock_author -bench=. -benchmem -v -cover ./...