package inception

import (
	"context"
	tf_core_framework "tensorflow/core/framework"
	pb "tensorflow_serving/apis"

	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
)

const (
	VERSION = "v0.0.1"
)

var (
	ModelName     = "mura"
	SignatureName = "predict_images"
)

func PredictRequest(b []byte) *pb.PredictRequest {
	request := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          ModelName,
			SignatureName: SignatureName,
			Version: &google_protobuf.Int64Value{
				Value: int64(1),
			},
		},
		Inputs: map[string]*tf_core_framework.TensorProto{
			"images": &tf_core_framework.TensorProto{
				Dtype: tf_core_framework.DataType_DT_STRING,
				TensorShape: &tf_core_framework.TensorShapeProto{
					Dim: []*tf_core_framework.TensorShapeProto_Dim{
						&tf_core_framework.TensorShapeProto_Dim{
							Size: int64(1),
						},
					},
				},
				StringVal: [][]byte{b},
			},
		},
	}
	return request
}

func Predict(Client pb.PredictionServiceClient, ctx context.Context, request *pb.PredictRequest) (*pb.PredictResponse, error) {
	resp, err := Client.Predict(ctx, request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
