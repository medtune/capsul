package main

import (
	"context"
	"fmt"
	"time"

	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

func main() {
	// Connect to client
	client, err := tfsclient.New("localhost:9002")
	if err != nil {
		panic(err)
	}

	// Prediction Request:
	meta := stdimpl.MuraMNV2
	req := pbreq.Status(meta)

	// Context with timeout 5seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ask for inception model status
	resp, err := client.Status(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
