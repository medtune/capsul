package capsul

import (
	"context"
	pb "tensorflow_serving/apis"

	"google.golang.org/grpc"
)

type Client struct {
	Address       string
	ModelName     string
	SignatureName string
	Version       int64
}

func (c *Client) Predict(ctx context.Context, request *pb.PredictRequest) (*pb.PredictResponse, error) {
	conn, err := grpc.Dial(c.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewPredictionServiceClient(conn)
	resp, err := client.Predict(ctx, request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *Client) PredictBytes(ctx context.Context, request *pb.PredictRequest) (*pb.PredictResponse, error) {
	conn, err := grpc.Dial(c.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewPredictionServiceClient(conn)
	resp, err := client.Predict(ctx, request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func New() {

}
