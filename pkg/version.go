package pkg

import (
	"fmt"
	"os"
)

const (
	VERSION = "v0.0.1"
)

func PrintAndExit() {
	fmt.Println(VERSION)
	os.Exit(0)
}
