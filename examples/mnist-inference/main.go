package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gocv.io/x/gocv"

	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
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

	// Prediction Request:
	meta := stdimpl.Mnist
	req := pbreq.PredictF32(meta, imgfloat)

	// run prediction
	resp, err := client.Predict(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
