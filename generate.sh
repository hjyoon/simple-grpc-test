#!/bin/sh

protoc \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  proto/helloworld.proto
