#!/bin/bash

# Build capsul dependencies: tensorflow framework & serving protos

# updates
apt-get update
apt update
# unzip
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

# tensorflow/serving
git clone -b r1.7 https://github.com/tensorflow/serving.git 
# tensorflow/tensorflow
git clone -b r1.7 https://github.com/tensorflow/tensorflow.git 

# protoc option
# -I tensorflow -I serving: paths where to search for imports
# --go_out: generate golang files using grpc plugin
# /go/src: where to generate files
PROTOC_OPTS='-I tensorflow -I serving --go_out=plugins=grpc:/go/src'

# compile proto files
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/apis/*.proto" 
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/config/*.proto" 
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/util/*.proto" 
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/sources/storage_path/*.proto" 
eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/framework/*.proto" 
eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/example/*.proto" 
eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/lib/core/*.proto" 
eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/protobuf/meta_graph.proto" 
eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/protobuf/saver.proto"

exit $?