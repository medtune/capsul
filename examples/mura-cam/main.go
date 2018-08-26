package main

import (
	"log"

	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

func main() {
	cli, err := tfsclient.Custom("localhost:8005/api", 1)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.Get(cli.Address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
