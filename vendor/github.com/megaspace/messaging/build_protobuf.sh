#!/bin/sh

protoc --go_out=plugins=grpc:. core/*.proto
protoc --go_out=plugins=grpc,Mcore/core.proto=github.com/megaspace/messaging/core:. stars/*.proto
