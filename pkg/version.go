package pkg

import (
	"fmt"
	"os"
)

const (
	VERSION = "v0.1.0"
)

func PrintAndExit() {
	fmt.Println(VERSION)
	os.Exit(0)
}
