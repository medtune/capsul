package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/medtune/capsul/pkg/request/inception"
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

	// Run prediction
	resp, err := client.Predict(ctx, inception.Default(b))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
