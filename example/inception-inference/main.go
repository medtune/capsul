package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/medtune/capsules/pkg/capsul"
	"github.com/medtune/capsules/pkg/request/inception"
)

func main() {
	b, err := ioutil.ReadFile("virus.jpeg")
	if err != nil {
		panic(err)
	}
	client, _ := capsul.New("localhost:9001")
	ctx := context.Background()
	resp, err := client.NewPredict(ctx, inception.Default(b))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
