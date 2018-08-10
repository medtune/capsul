#!/bin/bash

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