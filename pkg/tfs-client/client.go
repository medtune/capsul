package tfsclient

import (
	"context"
	"fmt"
	pb "tensorflow_serving/apis"

	"google.golang.org/grpc"
)

// Compile time check

var _ TFSClient = &Client{}

// Client wrap gRPC client
type Client struct {
	Address    string
	Connection *grpc.ClientConn
}

// New return a new client
// address is host:port example 127.0.0.1:9000
func New(address string) (*Client, error) {
	c := &Client{
		Address: address,
	}
	if err := c.Connect(); err != nil {
		return nil, fmt.Errorf("couldnt connect to capsul address: %v", err)
	}
	return c, nil
}

// Connect without closing
func (c *Client) Connect() error {
	conn, err := newInsecureConnection(c.Address)
	if err != nil {
		return err
	}
	c.Connection = conn
	return nil
}

// newInsecureConnection provide a grpc connection without tls
func newInsecureConnection(addr string) (*grpc.ClientConn, error) {
	return grpc.Dial(addr, grpc.WithInsecure())
}

// Close cached connection
func (c *Client) Close() error {
	return c.Connection.Close()
}

// Predict method
func (c *Client) Predict(ctx context.Context, request *pb.PredictRequest) (*pb.PredictResponse, error) {
	client := pb.NewPredictionServiceClient(c.Connection)
	return client.Predict(ctx, request)
}

// Status return model version status
func (c *Client) Status(ctx context.Context, request *pb.GetModelStatusRequest) (*pb.GetModelStatusResponse, error) {
	client := pb.NewModelServiceClient(c.Connection)
	return client.GetModelStatus(ctx, request)
}

// Cam method
func (c *Client) Cam(context.Context, *CamRequest) (*CamResponse, error) {
	return nil, fmt.Errorf("grad cam calculation isn't implemented in grpc clients")
}
