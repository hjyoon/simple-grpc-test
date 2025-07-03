# simple-grpc-test

This is a minimal example of a gRPC server and client in Go.

## Prerequisites

- `protoc` (Protocol Buffers compiler)
- Go plugins for code generation:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`
  
### Install plugins

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Quick Start

Follow these steps to run the example locally.

1. Generate Go code from proto files

```sh
./generate.sh
```

2. Run the server

```sh
go run ./cmd/server/main.go
```

3. Run the client

```sh
go run ./cmd/client/main.go
```
