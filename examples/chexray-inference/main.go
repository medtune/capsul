package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

type responsePP struct {
	Success bool   `json:"success"`
	Target  string `json:"target"`
	Out     string `json:"out"`
}

type imageData struct {
	data []float64
}

type requestPP struct {
	Target string `json:"target"`
}

func getFloatList(file string) (*responsePP, error) {
	url := "http://localhost:12030/process"

	jsonStr, err := json.Marshal(&requestPP{file})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	t := responsePP{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func main() {
	start := time.Now()
	// Read image file
	r, err := getFloatList("00000013_040.png")
	if err != nil {
		log.Fatalln(err)
	}

	var data [][][][][]float32
	err = json.Unmarshal([]byte(r.Out), &data)
	if err != nil {
		log.Fatalln(err)
	}

	var onelist []float32
	d := data[0][0]

	for _, i := range d {
		for _, j := range i {
			onelist = append(onelist, j[0], j[1], j[2])
		}
	}

	fmt.Println(len(onelist))

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
	req := pbreq.PredictFTest(meta, onelist)

	// Run prediction
	resp, err := client.Predict(ctx, req)
	if err != nil {
		panic(err)
	}

	l := resp.GetOutputs()["scores"].GetFloatVal()

	for _, i := range l {
		fmt.Printf("%.2f  ", i*float32(100))
	}
	fmt.Println(time.Since(start))
}
