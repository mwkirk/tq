SRC_ROOT := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
CMD_SERVER_DIR := $(SRC_ROOT)/cmd/server
CMD_WORKER_DIR := $(SRC_ROOT)/cmd/worker
BUILD_DIR := $(SRC_ROOT)/build
PROTO_DIR := $(SRC_ROOT)/pbuf

proto:
	cd $(PROTO_DIR); protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative tq.proto

clean:
	go clean

worker: proto
	cd $(CMD_WORKER_DIR); go build -o $(BUILD_DIR)/tq_worker

worker-linux: proto
	cd $(CMD_WORKER_DIR); GOOS=linux CGO_ENABLED=0 go build -o $(BUILD_DIR)/tq_worker

server: proto
	cd $(CMD_SERVER_DIR); go build -o $(BUILD_DIR)/tq_srv

server-linux: proto
	cd $(CMD_SERVER_DIR); GOOS=linux CGO_ENABLED=0 go build -o $(BUILD_DIR)/tq_srv

all: worker server

all-linux: worker-linux server-linux

.PHONY: clean worker worker-linux proto server server-linux docker
