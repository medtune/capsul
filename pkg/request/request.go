package request

import (
	tf_core_framework "tensorflow/core/framework"
	pb "tensorflow_serving/apis"

	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
)

func NewPredict(model, signature string, version int, inputData map[string]*tf_core_framework.TensorProto) *pb.PredictRequest {
	return &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          model,
			SignatureName: signature,
			Version: &google_protobuf.Int64Value{
				Value: int64(1),
			},
		},
		Inputs: inputData,
	}
}
