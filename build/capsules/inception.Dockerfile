FROM medtune/tensorflow-serving:bazel-cpu

RUN mkdir -p /models/inception

RUN curl -O http://download.tensorflow.org/models/image/imagenet/inception-v3-2016-03-01.tar.gz

RUN tar xzf inception-v3-2016-03-01.tar.gz

RUN bazel-bin/tensorflow_serving/example/inception_saved_model \
  --checkpoint_dir=inception-v3 \
  --output_dir=/models/inception

EXPOSE 10010

ENTRYPOINT [ "bazel-bin/tensorflow_serving/model_servers/tensorflow_model_server", \
    "--port=10010", \
    "--model_name=inception", \
    "--model_base_path=/models/inception"]