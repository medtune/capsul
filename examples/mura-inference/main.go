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
	client, err := tfsclient.New("35.242.219.80:10020")
	if err != nil {
		panic(err)
	}

	// Context with timeout 5seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Prediction Request:
	meta := stdimpl.MuraMNV2
	req := pbreq.Predict(meta, b)

	start := time.Now()
	// Run prediction
	resp, err := client.Predict(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Since(start))

	fmt.Println(resp)
}
