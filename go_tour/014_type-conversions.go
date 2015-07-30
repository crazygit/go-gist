// 与 C 不同的是 Go 的在不同类型之间的项目赋值时需要显式转换

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64((x*x + y*y)))
	var z int = int(f)
	fmt.Println(x, y, z)

}
