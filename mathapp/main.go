package main

import (
	"../mymath"
	"fmt"
)

func main() {
	fmt.Printf("Hello world. Sqrt(2) =%v \n", mymath.Sqrt(2))
	a := [...]int{1, 2, 3, 4}
	fmt.Println(a) // [1 2 3 4]

	changearry(a)
	fmt.Println(a) //[1 2 3 4]

	var aslice []int
	fmt.Println(len(aslice)) // 0
	fmt.Println(cap(aslice)) // 0

	bslice := []int{1, 2, 3}
	fmt.Println(len(bslice)) // 3
	fmt.Println(cap(bslice)) // 3

	array_a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice_a := array_a[2:4:7]
	fmt.Println(len(slice_a)) // 2
	fmt.Println(cap(slice_a)) // 5
	fmt.Println(slice_a)      // 3 4
	slice_b := array_a[2:5:7]
	slice_c := array_a[2:6:8]
	fmt.Printf("slice c is %v \n", slice_c)
	slice_b = append(slice_b, 1, 2, 3)
	fmt.Println(slice_a) // 3 4

	fmt.Println()
	fmt.Println(slice_b) // 3 4
	fmt.Printf("slice c is %v \n", slice_c)
}

func changearry(arr [4]int) {
	arr[2] = 7
}
