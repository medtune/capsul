package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getppid(), os.Getpid())
	fmt.Println(os.Getuid())
	fmt.Println(os.Hostname())
	fmt.Println(os.Setenv("HELLO", "EZ"))
}
