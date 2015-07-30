//基本的 for 循环除了没有了 ( ) 之外（甚至强制不能使用它们）
package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
