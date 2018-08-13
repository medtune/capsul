package request

import (
	tf_core_framework "tensorflow/core/framework"
	pb "tensorflow_serving/apis"

	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
)

// Predict request
func Predict(model, signature string, version int, inputData map[string]*tf_core_framework.TensorProto) *pb.PredictRequest {
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

// GetStatus request
func Status(model string, version int) *pb.GetModelStatusRequest {
	return &pb.GetModelStatusRequest{
		ModelSpec: &pb.ModelSpec{
			Name: model,
			Version: &google_protobuf.Int64Value{
				Value: int64(version),
			},
		},
	}
}
