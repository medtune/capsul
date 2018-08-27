package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/medtune/capsul/pkg/request/mura"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

func main() {
	// Read image file
	b, err := ioutil.ReadFile("test/testdata/image2.png")
	if err != nil {
		panic(err)
	}

	// Connection to tf server
	client, err := tfsclient.New("localhost:9002")
	if err != nil {
		panic(err)
	}

	// Context with timeout 5seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Run prediction
	resp, err := client.Predict(ctx, mura.Default(b))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
