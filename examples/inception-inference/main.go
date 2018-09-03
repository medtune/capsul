package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

func main() {
	// Read image file
	b, err := ioutil.ReadFile("test/testdata/inception_cheetah.jpeg")
	if err != nil {
		panic(err)
	}

	// Connection to tf server
	client, err := tfsclient.New("localhost:9001")
	if err != nil {
		panic(err)
	}

	// timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Prediction Request:
	meta := stdimpl.Inception
	req := pbreq.Predict(meta, b)

	// Run prediction
	resp, err := client.Predict(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
