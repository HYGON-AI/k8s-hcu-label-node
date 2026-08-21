#!/bin/bash
# SPDX-License-Identifier: Apache-2.0
# Copyright 2026 Hygon Information Technology Co., Ltd.

# Build the binary
export GOPROXY=https://goproxy.cn
export CGO_ENABLED=1
go mod tidy
go build -o hcu-label-node cmd/main.go

# Build the docker image
docker build --target label-node -t harbor.sourcefind.cn:5443/hcu/admin/base/hcu-label-node:"$(git describe --tags --abbrev=0)" .
docker save -o hcu-label-node-"$(git describe --tags --abbrev=0)".tar harbor.sourcefind.cn:5443/hcu/admin/base/hcu-label-node:"$(git describe --tags --abbrev=0)"