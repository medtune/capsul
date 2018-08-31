FROM golang:1.10

# Update system (ubuntu)
RUN apt-get update && apt update

# Add project to default GOPATH
ADD cmd /go/src/github.com/medtune/capsul/cmd
ADD pkg /go/src/github.com/medtune/capsul/pkg
ADD hack /go/src/github.com/medtune/capsul/hack
ADD plugins /go/src/github.com/medtune/capsul/plugins
ADD csflask /go/src/github.com/medtune/capsul/csflask
ADD examples /go/src/github.com/medtune/capsul/examples

# install unzip
RUN apt install unzip -y

# Download protoc binaries
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip && \
    unzip ./protoc-3.6.1-linux-x86_64.zip -d protoc3

# Move protoc binaries
RUN mv protoc3/bin/* /usr/local/bin/ && \
    mv protoc3/include/* /usr/local/include/

# Install golang protoc & grpc
RUN go get github.com/golang/protobuf/ptypes/wrappers
RUN go get google.golang.org/grpc
RUN go get github.com/golang/protobuf/protoc-gen-go

# Git clone tensorflow/tensorflow and tensorflow/serving version 1.7 
#NOTE: earlier version (>1.7) fail to compile ?
RUN git clone -b r1.7 https://github.com/tensorflow/serving.git 
RUN git clone -b r1.7 https://github.com/tensorflow/tensorflow.git 

# Compile proto files using gRPC plugin 
RUN PROTOC_OPTS='-I tensorflow -I serving --go_out=plugins=grpc:src' && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/apis/*.proto" && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/config/*.proto" && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/util/*.proto" && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/sources/storage_path/*.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/framework/*.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/example/*.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/lib/core/*.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/protobuf/meta_graph.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/protobuf/saver.proto"

#TODO: Install gocv: opencv bindings

# Set work dir
WORKDIR /go/src/github.com/medtune/capsul

#RUN go install ./example/inception-inference
#RUN go install ./example/model-status
#RUN go install ./example/mnist-inference/main.go