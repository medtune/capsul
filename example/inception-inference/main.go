package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/medtune/capsul/pkg/request/inception"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

func main() {
	// Read image file
	b, err := ioutil.ReadFile("testdata/inception_cheetah.jpeg")
	if err != nil {
		panic(err)
	}

	// Connection to tf server
	client, err := tfsclient.New("localhost:9001")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	// Run prediction
	resp, err := client.Predict(ctx, inception.Default(b))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
