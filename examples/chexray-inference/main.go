package main

import (
	"bytes"
	"context"
	"fmt"
	"image/png"
	"io/ioutil"
	"time"

	"github.com/anthonynsimon/bild/transform"
	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
	"gocv.io/x/gocv"
)

// Convert images bytes to float32 array
// Need gocv
func bytesToFloat32(ib []byte) ([]float32, error) {
	im, _ := png.Decode(bytes.NewReader(ib))
	im = transform.Resize(im, 224, 224, transform.Linear)
	m, _ := gocv.ImageToMatRGB(im)
	mat := gocv.NewMat()
	m.ConvertTo(&mat, gocv.MatTypeCV32F)

	imgfloat := make([]float32, 0, 0)
	for i := 0; i < 224; i++ {
		for j := 0; j < 224; j++ {
			a, b, c, alpha := im.At(i, j).RGBA()
			var alp = float32(alpha)
			imgfloat = append(imgfloat,
				(float32(a)/alp-0.229)/0.485,
				(float32(b)/alp-0.224)/0.456,
				(float32(c)/alp-0.225)/0.406)
		}
	}
	fmt.Println(len(imgfloat))
	return imgfloat, nil
}

func main() {
	// Read image file
	b, err := ioutil.ReadFile("/Volumes/IAL-CLOUD/Chest-X-Ray/00000013_040.png")
	if err != nil {
		panic(err)
	}

	// Connection to tf server
	client, err := tfsclient.New("localhost:10031")
	if err != nil {
		panic(err)
	}

	// timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Prediction Request:
	meta := stdimpl.ChexrayDN121
	f, _ := bytesToFloat32(b)
	req := pbreq.PredictFTest(meta, f)

	// Run prediction
	resp, err := client.Predict(ctx, req)
	if err != nil {
		panic(err)
	}

	l := resp.GetOutputs()["scores"].GetFloatVal()

	for _, i := range l {
		fmt.Printf("%.2f  ", i*float32(100))
	}
}
