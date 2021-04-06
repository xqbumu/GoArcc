#!/usr/bin/env sh
# put the protoc path of your local
protoc -I .  --proto_path=../../../../protos  --grpc-gateway_out=./  --grpc-gateway_opt logtostderr=true --go_out=./  --go-grpc_out=./ --graphql_out=./   ./*.proto
go fmt ./...
go mod tidy
