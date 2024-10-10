## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## tidy: run go mod tidy && go fmt ./...
.PHONY: tidy
tidy:
	go mod tidy
	go fmt ./...

## generate: run go generate && go fmt
.PHONY: generate
generate:
	go generate ./...
	go fmt ./...
	
## vet: vet files
.PHONY: vet
vet:
	go vet ./...

## test: run tests
.PHONY: test
test:
	op run --env-file='./.env' -- go test ./...

