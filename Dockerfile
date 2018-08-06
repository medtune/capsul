FROM golang:1.10

WORKDIR /go/src

ADD . github.com/medtune/capsules/

RUN apt-get update && apt update

RUN apt install unzip -y

RUN curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip && \
    ls -la && unzip ./protoc-3.6.1-linux-x86_64.zip -d protoc3

RUN mv protoc3/bin/* /usr/local/bin/ && \
    mv protoc3/include/* /usr/local/include/

RUN go get -u github.com/golang/protobuf/ptypes/wrappers
RUN go get -u google.golang.org/grpc/...
RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN git clone https://github.com/tensorflow/serving.git
RUN git clone https://github.com/tensorflow/tensorflow.git

RUN mkdir -p vendor && PROTOC_OPTS='-I tensorflow -I serving --go_out=plugins=grpc:vendor' && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/apis/*.proto" && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/config/*.proto" && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/util/*.proto" && \
    eval "protoc $PROTOC_OPTS serving/tensorflow_serving/sources/storage_path/*.proto"

RUN export PROTOC_OPTS='-I tensorflow -I serving --go_out=plugins=grpc:vendor' && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/framework/*.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/example/*.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/lib/core/*.proto" && echo hello && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/protobuf/meta_graph.proto" && \
    eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/protobuf/saver.proto"


RUN ls -la

RUN go get -v -u github.com/medtune/capsules/example/inception-inference
RUN go build github.com/medtune/capsules/example/inception-inference/main.go

RUN ls -la

