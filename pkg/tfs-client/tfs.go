package tfsclient

import (
	"context"
	pb "tensorflow_serving/apis"
)

// TFSClient .
type TFSClient interface {
	Status(context.Context, *pb.GetModelStatusRequest) (*pb.GetModelStatusResponse, error)
	Predict(context.Context, *pb.PredictRequest) (*pb.PredictResponse, error)
	Cam(context.Context, *pb.PredictRequest) (*pb.PredictResponse, error)
}
