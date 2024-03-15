SRC_ROOT := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
CMD_CLI_DIR := $(SRC_ROOT)/cmd/cli
CMD_SERVER_DIR := $(SRC_ROOT)/cmd/server
CMD_WORKER_DIR := $(SRC_ROOT)/cmd/worker
BUILD_DIR := $(SRC_ROOT)/build
PROTO_DIR := $(SRC_ROOT)/pb

.DEFAULT_GOAL := all

clean:
	go clean

proto:
	cd $(PROTO_DIR); protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto

cli: proto
	cd $(CMD_CLI_DIR); go build -o $(BUILD_DIR)/tq

cli-linux: proto
	cd $(CMD_CLI_DIR); GOOS=linux CGO_ENABLED=0 go build -o $(BUILD_DIR)/tq-${GOOS}

worker: proto
	cd $(CMD_WORKER_DIR); go build -o $(BUILD_DIR)/tq_worker

worker-linux: proto
	cd $(CMD_WORKER_DIR); GOOS=linux CGO_ENABLED=0 go build -o $(BUILD_DIR)/tq_worker-${GOOS}

server: proto
	cd $(CMD_SERVER_DIR); go build -o $(BUILD_DIR)/tq_srv

server-linux: proto
	cd $(CMD_SERVER_DIR); GOOS=linux CGO_ENABLED=0 go build -o $(BUILD_DIR)/tq_srv-${GOOS}

test:
	go test ./...

all: cli worker server

all-linux: cli-linux worker-linux server-linux

.PHONY: clean cli worker worker-linux proto server server-linux docker
