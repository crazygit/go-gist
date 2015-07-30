package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
	fmt.Println()

	var g int
	go func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
		}
		g = s
	}(1000)
	fmt.Println(s)
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
