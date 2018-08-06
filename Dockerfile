FROM golang:1.10

# Update system (ubuntu)
RUN apt-get update && apt update

# Add project to default GOPATH
ADD . /go/src/github.com/medtune/capsules/

# install unzip
RUN apt install unzip -y

# Download protoc binaries
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip && \
    unzip ./protoc-3.6.1-linux-x86_64.zip -d protoc3

# Move protoc binaries
RUN mv protoc3/bin/* /usr/local/bin/ && \
    mv protoc3/include/* /usr/local/include/

# Install golang protoc & grpc
RUN go get -u github.com/golang/protobuf/ptypes/wrappers
RUN go get -u google.golang.org/grpc/...
RUN go get -u github.com/golang/protobuf/protoc-gen-go

# Git clone tensorflow/tensorflow and tensorflow/serving
RUN git clone https://github.com/tensorflow/serving.git 
RUN git clone https://github.com/tensorflow/tensorflow.git 

RUN ls -la /go/src

# Compile proto files using gRPC plugin 
RUN PROTOC_OPTS='-I serving --go_out=plugins=grpc:/go/src' && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/apis/*.proto" && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/config/*.proto" && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/util/*.proto" && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/sources/storage_path/*.proto" && \

RUN ls -la /go/src

RUN PROTOC_OPTS='-I tensorflow --go_out=plugins=grpc:/go/src' && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/framework/*.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/example/*.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/lib/core/*.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/protobuf/meta_graph.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/protobuf/saver.proto"

WORKDIR /go/src

RUN ls -la

RUN go build github.com/medtune/capsules/example/inception-inference/main.go

RUN ls -la
