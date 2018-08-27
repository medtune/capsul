package main

import (
	"context"
	"fmt"

	"github.com/medtune/capsul/pkg/request"
	"github.com/medtune/capsul/pkg/request/mura"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

func main() {
	// Connect to client
	client, err := tfsclient.New("localhost:9002")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// Ask for inception model status
	resp, err := client.Status(ctx, request.Status(mura.Model, 1))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
