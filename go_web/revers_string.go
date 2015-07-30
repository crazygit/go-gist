package main

import (
	"fmt"
	"sort"
)

func a() {
	i := 0
	defer fmt.Printf("defer i:%d\n", i)
	i++
	fmt.Printf("i:%d\n", i)
	return
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	fmt.Println(c())
	a()
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	j := 0
LOOP:
	fmt.Print(j)
	if j++; j < 10 {
		goto LOOP
	}

	fmt.Println()
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range a {
		fmt.Print(v)
	}
	fmt.Println()
	s := "abcedfg"
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Print(s[i : i+1])
	}
	fmt.Println()
	Revert("abcdefg")
	fmt.Println(RevertString("abcdefdxs"))
	fmt.Println(sort.Reverse(sort.StringSlice{"abcdefg"}))
}

func RevertString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}
func Revert(s string) {
	var result string
	for _, val := range s {
		fmt.Println(string(val))
		result = string(val) + result
	}
	fmt.Println(result)
}
