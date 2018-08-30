package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/medtune/capsul/pkg/pbreq/stdimpl"

	"github.com/medtune/capsul/pkg/pbreq"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

// Get model status

var (
	model string
	host  string
	port  int
)

func init() {
	flag.StringVar(&model, "model", "mnist", "model name")
	flag.StringVar(&host, "host", "localhost", "serving host")
	flag.IntVar(&port, "port", 9000, "serving port")
	flag.Parse()
}

func main() {
	// default will be 'localhost:9000'
	endpoint := fmt.Sprintf("%s:%d", host, port)

	// Connect to client
	client, err := tfsclient.New(endpoint)
	if err != nil {
		panic(err)
	}

	// Context with timeout 5seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	meta := stdimpl.Map[model]

	// Ask for inception model status
	resp, err := client.Status(ctx, pbreq.Status(meta))
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
