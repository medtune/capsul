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
	b, err := ioutil.ReadFile("test/testdata/mura_0.png")
	if err != nil {
		panic(err)
	}

	// Connection to tf server
	client, err := tfsclient.New("localhost:10021")
	if err != nil {
		panic(err)
	}

	// Context with timeout 5seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Prediction Request:
	meta := stdimpl.MuraIRNV2
	req := pbreq.Predict(meta, b)

	// Run prediction
	resp, err := client.Predict(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
