export GO111MODULE=on

GOBIN ?= $(shell go env GOPATH)/bin
DEP_PROGS=\
	$(GOBIN)/golangci-lint\
	$(GOBIN)/kratos\
	$(GOBIN)/wire\
	$(GOBIN)/swagger\
	$(GOBIN)/protoc-gen-validate\
	$(GOBIN)/protoc-gen-go-errors\
	$(GOBIN)/protoc-gen-go-http\
	$(GOBIN)/protoc-gen-grpc-gateway\
	$(GOBIN)/protoc-gen-openapiv2\
	$(GOBIN)/protoc-gen-go-grpc\
	$(GOBIN)/protoc-gen-go

.PHONY: all
all: clean build

.PHONY: clean
clean:
	@echo $(ROOT_DIR)
	rm -f $(DEP_PROGS)

.PHONY: build
build: $(DEP_PROGS)

deps:
	go mod tidy
.PHONY: deps

$(GOBIN)/golangci-lint: deps
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

$(GOBIN)/kratos: deps
	go install github.com/go-kratos/kratos/cmd/kratos/v2

$(GOBIN)/wire: deps
	go install github.com/google/wire/cmd/wire

$(GOBIN)/swagger: deps
	go install github.com/go-swagger/go-swagger/cmd/swagger

$(GOBIN)/protoc-gen-validate: deps
	go install github.com/envoyproxy/protoc-gen-validate

$(GOBIN)/protoc-gen-go-errors: deps
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2

$(GOBIN)/protoc-gen-go-http: deps
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2

$(GOBIN)/protoc-gen-grpc-gateway: deps
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

$(GOBIN)/protoc-gen-openapiv2: deps
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

$(GOBIN)/protoc-gen-go-grpc: deps
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

$(GOBIN)/protoc-gen-go: deps
	go install google.golang.org/protobuf/cmd/protoc-gen-go