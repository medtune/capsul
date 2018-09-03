package pkg

import (
	"fmt"
	"os"
)

const (
	VERSION = "v0.0.2"
)

// PrintAndExit .
func PrintAndExit() {
	fmt.Println(VERSION)
	os.Exit(0)
}
