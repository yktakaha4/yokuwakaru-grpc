#!/usr/bin/make -f

.PHONY: server
server:
	cd ./hello-grpc/server/ && go run server.go ${port}

.PHONY: client
client:
	cd ./hello-grpc/client/ && go run client.go ${message}

.PHONY: client-round-robin
client-round-robin:
	cd ./hello-grpc/client/ && go run client-round-robin.go ${message}

.PHONY: proto
proto:
	cd ./hello-grpc/ && protoc greeter.proto --go_out=plugins=grpc:.
