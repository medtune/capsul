package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/medtune/capsules/pkg/capsul"
	"github.com/medtune/capsules/pkg/request/mnist"
	"gocv.io/x/gocv"
)

func main() {

	b, err := ioutil.ReadFile("1.png")
	if err != nil {
		panic(err)
	}
	client, _ := capsul.New("localhost:9000")
	ctx := context.Background()
	matb, err := gocv.IMDecode(b, -1)
	if err != nil {
		panic(fmt.Errorf("not an image %v", err))
	}

	mat := gocv.NewMat()
	matb.ConvertTo(&mat, gocv.MatTypeCV32F)

	imgfloat := make([]float32, 0, 0)

	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			//fmt.Printf("value: %v %v: %v %v\n", i, j, matb.GetFloatAt(i, j), mat.GetFloatAt(i, j))
			imgfloat = append(imgfloat, mat.GetFloatAt(i, j))
		}
	}

	resp, err := client.NewPredict(ctx, mnist.Default(imgfloat))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
