package main

import (
	"fmt"
)

func fibonacci() func() int {
	init := 0
	next := 1
	return func() int {
		third := init + next
		init = next
		next = third
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
}
