GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto: server-proto proxy-proto

.PHONY: server-proto
server-proto:
	protoc -I/usr/local/include -I. \
	  -I${GOPATH}/src \
	  -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	  --go_out=plugins=grpc:. \
	  graderpb/grader.proto

.PHONY: proxy-proto
proxy-proto:
	protoc -I/usr/local/include -I. \
      -I${GOPATH}/src \
      -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --grpc-gateway_out=logtostderr=true:. \
      graderpb/grader.proto

.PHONY: build
build: proto
	go build -o bin/graderd cmd/graderd/main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t graderd:latest