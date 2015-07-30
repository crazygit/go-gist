package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(strings.HasPrefix("hello", "he"))
	fmt.Println(strings.HasSuffix("hello", "lo"))
	fmt.Println(strings.Contains("hello", "el"))
	fmt.Println(strings.Count("hello", "l"))
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())

	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.Compiler)

	var orig string = "123"
	var newS string
	fmt.Printf("The size of ints is : %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %q\n", newS)

	k := 4
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
