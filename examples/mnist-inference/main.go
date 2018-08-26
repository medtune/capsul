package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/medtune/capsul/pkg/request/mnist"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"

	"gocv.io/x/gocv"
)

func main() {
	// Read file
	b, err := ioutil.ReadFile("test/testdata/mnist_1.png")
	if err != nil {
		panic(err)
	}

	// Connect to tf server
	client, err := tfsclient.New("localhost:9000")
	if err != nil {
		panic(err)
	}

	// Image to float32 array using gocv

	// Decode image
	matb, err := gocv.IMDecode(b, -1)
	if err != nil {
		panic(fmt.Errorf("not an image %v", err))
	}

	// Convert matrix
	mat := gocv.NewMat()
	matb.ConvertTo(&mat, gocv.MatTypeCV32F)

	// matrix to array
	imgfloat := make([]float32, 0, 0)
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			imgfloat = append(imgfloat, mat.GetFloatAt(i, j))
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// run prediction
	resp, err := client.Predict(ctx, mnist.Default(imgfloat))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
