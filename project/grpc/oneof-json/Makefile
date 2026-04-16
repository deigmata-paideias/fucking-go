GOBIN := $(shell go env GOBIN)
GOPATH_BIN := $(shell go env GOPATH)/bin
PROTOC_GEN_GO := $(if $(GOBIN),$(GOBIN)/protoc-gen-go,$(GOPATH_BIN)/protoc-gen-go)

.PHONY: tools generate tidy run verify

tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11

generate: tools
	protoc --plugin=protoc-gen-go=$(PROTOC_GEN_GO) --go_out=. --go_opt=paths=source_relative proto/hello.proto

tidy:
	go mod tidy

run:
	go run .

verify: generate tidy run
