package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/medtune/capsules/pkg/request/mnist"
	tfsclient "github.com/medtune/capsules/pkg/tfs-client"

	"gocv.io/x/gocv"
)

func main() {
	// Read file
	b, err := ioutil.ReadFile("1.png")
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

	ctx := context.Background()

	// run prediction
	resp, err := client.Predict(ctx, mnist.Default(imgfloat))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
