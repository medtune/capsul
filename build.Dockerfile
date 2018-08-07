FROM golang:1.10

WORKDIR /go/src/github.com/medtune/capsul

ADD . .

RUN chmod +x build.sh && ./build.sh

# Set work dir
WORKDIR /go/src/github.com/medtune/capsul

RUN go install ./example/inception-inference
#RUN go install ./example/mnist-inference

