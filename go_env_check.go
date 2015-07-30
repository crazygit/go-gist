package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("GOROOT: ", os.Getenv("GOROOT"))
	fmt.Println("GOPATH: ", os.Getenv("GOPATH"))
}
