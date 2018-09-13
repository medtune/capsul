package tfsclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	pb "tensorflow_serving/apis"
	"time"

	"encoding/json"
)

const (
	baseEP    = "/"
	camEP     = "/cam"
	listEP    = "/list"
	statusEP  = "/status"
	predictEP = "/predict"
	processEP = "/process"
)

// Compile check
var _ TFSClient = &RestClient{}

// RestClient .
type RestClient struct {
	*http.Client
	Address string
}

// NewRest .
func NewRest(address string, timeout int) (*RestClient, error) {
	addr, err := url.Parse("http://" + address)
	if err != nil {
		return nil, err
	}

	tt := time.Duration(timeout) * time.Second

	// Create client
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

// CamRequest .
type CamRequest struct {
	Target string `json:"target"`
	Dest   string `json:"destination"`
	Force  bool   `json:"force"`
}

// CamResponse .
type CamResponse struct {
	Success bool     `json:"success"`
	Target  string   `json:"target"`
	Dest    string   `json:"destination"`
	Errors  []string `json:"errors"`
}

// Cam .
func (rc *RestClient) Cam(ctx context.Context, r *CamRequest) (*CamResponse, error) {
	url := rc.Address + camEP

	jsonStr, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	camResp := CamResponse{}
	err = json.Unmarshal(body, &camResp)
	if err != nil {
		return nil, err
	}

	return &camResp, nil
}

type statusResponse struct {
	Success bool   `json:"success"`
	Status  string `json:"status"`
	Version int64  `json:"version"`
}

// Status .
func (rc *RestClient) Status(ctx context.Context, r *pb.GetModelStatusRequest) (*pb.GetModelStatusResponse, error) {
	url := rc.Address + statusEP

	if r == nil {
		r = &pb.GetModelStatusRequest{}
	}

	jsonStr, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	statusResp := statusResponse{}
	err = json.Unmarshal(body, &statusResp)
	if err != nil {
		return nil, err
	}

	if !statusResp.Success {
		return &pb.GetModelStatusResponse{
			ModelVersionStatus: []*pb.ModelVersionStatus{
				&pb.ModelVersionStatus{
					Version: statusResp.Version,
					State:   pb.ModelVersionStatus_UNKNOWN,
				},
			},
		}, nil
	}

	return &pb.GetModelStatusResponse{
		ModelVersionStatus: []*pb.ModelVersionStatus{
			&pb.ModelVersionStatus{
				Version: statusResp.Version,
				State:   pb.ModelVersionStatus_AVAILABLE,
			},
		},
	}, nil

}

// ProcessRequest .
type ProcessRequest struct {
	Target string `json:"target"`
}

// ProcessResponse .
type ProcessResponse struct {
	Success bool     `json:"success"`
	Out     string   `json:"out"`
	Target  string   `json:"target"`
	Errors  []string `json:"errors"`
}

// Process .
func (rc *RestClient) Process(ctx context.Context, preq *ProcessRequest) (*ProcessResponse, error) {
	url := rc.Address + processEP

	// Marshal request body
	jsonStr, err := json.Marshal(&ProcessRequest{preq.Target})
	if err != nil {
		return nil, err
	}

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// dont forget to close body
	defer r.Body.Close()

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Parse response
	t := ProcessResponse{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
