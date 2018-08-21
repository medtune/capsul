package tfsclient

import (
	"context"
	"net"
	"net/http"
	"net/url"
	pb "tensorflow_serving/apis"
	"time"
)

const (
	baseEP    = "/"
	predictEP = "/predict"
	camEP     = "/cam"
	statusEP  = "/status"
)

var _ TFSClient = &RestClient{}

// RestClient .
type RestClient struct {
	*http.Client
	Address string
}

// Custom .
func Custom(address string, timeout int) (*RestClient, error) {
	tt := time.Duration(timeout) * time.Second

	addr, err := url.Parse("http://" + address)
	if err != nil {
		return nil, err
	}
	// Client
	client := &RestClient{
		Address: addr.String(),
		Client: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: tt,
				}).Dial,
			},
			Timeout: tt,
		},
	}
	return client, nil
}

// Predict .
func (rc *RestClient) Predict(context.Context, *pb.PredictRequest) (*pb.PredictResponse, error) {
	return nil, nil
}

// Cam .
func (rc *RestClient) Cam(context.Context, *pb.PredictRequest) (*pb.PredictResponse, error) {
	return nil, nil
}

// Status .
func (rc *RestClient) Status(context.Context, *pb.GetModelStatusRequest) (*pb.GetModelStatusResponse, error) {
	return nil, nil
}
