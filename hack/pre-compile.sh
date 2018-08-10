#!/bin/bash

#unzip
apt install -y unzip 
#curl
apt install -y curl

# protoc 3.6.1 
curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
unzip ./protoc-3.6.1-linux-x86_64.zip -d protoc3
mv protoc3/bin/* /usr/local/bin/
mv protoc3/include/* /usr/local/include/

# protobuf wrappers
go get github.com/golang/protobuf/ptypes/wrappers
# gRPC golang
go get google.golang.org/grpc
# protoc go plugin
go get github.com/golang/protobuf/protoc-gen-go
