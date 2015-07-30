package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello, world")
	fmt.Println(os.Getenv("IP"))
	fmt.Println(os.Getenv("PORT"))
}
