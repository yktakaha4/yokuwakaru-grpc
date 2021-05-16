#!/usr/bin/make -f

.PHONY: server
server:
	cd ./hello-grpc/server/ && go run server.go

.PHONY: client
client:
	cd ./hello-grpc/client/ && go run client.go ${message}
