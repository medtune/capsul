package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/medtune/capsul/pkg/pbreq"
	"github.com/medtune/capsul/pkg/pbreq/stdimpl"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

func main() {
	// Read image file
	b, err := ioutil.ReadFile("test/testdata/mura_0.png")
	if err != nil {
		panic(err)
	}

	// Context with timeout 5seconds
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Prediction Request:
	meta := stdimpl.MuraMNV2
	req := pbreq.Predict(meta, b)

	var wg sync.WaitGroup
	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go func() {
			client, _ := tfsclient.New("35.242.219.80:10020")
			start := time.Now()
			// Run prediction
			_, err := client.Predict(ctx, req)
			log.Println(time.Since(start), err)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("done")
}
