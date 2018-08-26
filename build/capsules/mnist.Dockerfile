FROM medtune/tensorflow-serving:bazel-cpu

RUN mkdir -p /models/mnist

RUN bazel-bin/tensorflow_serving/example/mnist_saved_model /models/mnist

ENTRYPOINT [ "bazel-bin/tensorflow_serving/model_servers/tensorflow_model_server", \
    "--port=9000", \
    "--model_name=mnist", \
    "--model_base_path=/models/mnist"]