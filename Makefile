PATH := $(CURDIR)/.local/bin:$(CURDIR)/.local/go/bin:$(PATH)
PROTOC_URL := https://github.com/protocolbuffers/protobuf/releases/download/v29.0/protoc-29.0-osx-aarch_64.zip
GO_INSTALL := GOBIN=$(CURDIR)/.local/go/bin go install

.PHONY: init
init:
	curl -o .local/tmp/protoc.zip -fsSL $(PROTOC_URL)
	unzip -q .local/tmp/protoc.zip -d .local
	$(GO_INSTALL) google.golang.org/protobuf/cmd/protoc-gen-go@latest

.PHONY: generate
generate:
	.local/bin/protoc --go_out=./generated --go_opt=paths=source_relative $(shell find ./pb -name '*.proto')

.PHONY: test
test:
	go test ./...
