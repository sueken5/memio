#!/bin/bash

set -euxo pipefail

protoc --go_out=plugins=grpc:./server ./apis/v1/memio.proto
protoc --js_out=import_style=commonjs:./web/src --grpc-web_out=import_style=typescript,mode=grpcwebtext:./web/src ./apis/v1/memio.proto

