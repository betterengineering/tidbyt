#!/bin/bash

# Notes: this script should be run in the project root. You will need to set TIDBYT_DEVICE_ID and TIDBYT_AUTH_TOKEN
# as environment variables for the config to setup properly. Finally, this script assumes you have protoc already
# installed and that you have a standard go installation.

# Fetch most recent
go run cmd/reflect-to-proto/main.go

# Fetch imported proto definitions.
GO111MODULE=off go get github.com/grpc-ecosystem/grpc-gateway/...
GO111MODULE=off go get github.com/googleapis/googleapis/...

# Generate client.
protoc \
	--go_out=plugins=grpc:api \
	-I api/ \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	-I ${GOPATH}/src/github.com/googleapis/googleapis \
	api/public-api/proto/public_api.proto

