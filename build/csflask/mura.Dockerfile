FROM medtune/tensorflow-serving:bazel-cpu

RUN mkdir -p /models/mnist

COPY murackpt .

