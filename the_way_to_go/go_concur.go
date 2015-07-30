package main

import (
	"fmt"
)

var complete chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	complete <- 0
}

func main() {
	go loop()
	loop()
	<-complete
}
