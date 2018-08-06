package capsul

import (
	"context"
	pb "tensorflow_serving/apis"

	"google.golang.org/grpc"
)

// Client wrap gRPC client
type Client struct {
	Address    string
	CachedConn *grpc.ClientConn
}

func New(address string) (*Client, error) {
	return &Client{
		Address: address,
	}, nil
}

// NewPredict send a predict request to the designed server
// using a new fresh connection and closing it later
func (c *Client) NewPredict(ctx context.Context, request *pb.PredictRequest) (*pb.PredictResponse, error) {
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

// Predict method
func (c *Client) Predict(ctx context.Context, request *pb.PredictRequest) (*pb.PredictResponse, error) {
	if c.CachedConn == nil {
		c.Connect()
	}
	client := pb.NewPredictionServiceClient(c.CachedConn)
	resp, err := client.Predict(ctx, request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Connect without closing
func (c *Client) Connect() error {
	conn, err := grpc.Dial(c.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.CachedConn = conn
	return nil
}

// Close cached conn
func (c *Client) Close() error {
	err := c.CachedConn.Close()
	c.CachedConn = nil
	return err
}

// HealthCheck not implemented yet
func (c *Client) HealthCheck(ctx context.Context) {

}
