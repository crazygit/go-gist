//基于此可以省略分号：C 的 while 在 Go 中叫做 for 。
package main

import "fmt"

func main() {

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
