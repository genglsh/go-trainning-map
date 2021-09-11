// +build tools

package tools

// Package tools is used to track binary dependencies with go modules
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
import (
	// linter tools
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"

	// kratos
	_ "github.com/go-kratos/kratos/cmd/kratos/v2"

	// wire
	_ "github.com/google/wire/cmd/wire"

	// swagger
	_ "github.com/go-swagger/go-swagger/cmd/swagger"

	// protoc generate tools and plugin
	_ "github.com/envoyproxy/protoc-gen-validate"
	_ "github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2"
	_ "github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
