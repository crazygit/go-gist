package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
	C    int
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	//	xx := &Vertex{3,4, 7}
	v := Vertex{3, 4, 7}
	v.Scale(5)
	//	v.C = 8
	//	test(v)
	//	test1(xx)
	fmt.Println(v, v.Abs())
	//	fmt.Println(xx, xx.Abs())
}

func test(tt Vertex) {
	tt.C = 100
}

func test1(tt *Vertex) {
	tt.C = 200
}

//package main
//
////学习指针用法
//
//import (
//	"fmt"
//)
//
//func main() {
//	var i int; // i 的类型是int型
//	var p *int; // p 的类型是[int型的指针]
//	i = 1; // i 的值为 1;
//	p = &i; // p 的值为 [i的地址]
//	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p);
//	*p = 2; // *p 的值为 [[i的地址]的指针](其实就是i嘛),这行代码也就等价于 i = 2
//	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p);
//	i = 3; // 验证我的想法
//	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p);
//}

//package main

//学习函数参数的用法
//
//import (
//    "fmt"
//)
//
//type abc struct {
//     v int;
//}
//
//func (a abc) aaaa (){
//    a.v = 1;
//    fmt.Printf("1:%d\n",a.v);  //1
//}
//
//func (a *abc) bbbb (){
//    fmt.Printf("2:%d\n",a.v); // 0
//    a.v = 2;
//    fmt.Printf("3:%d\n",a.v); // 2
//}
//
//func (a *abc) cccc(){
//    fmt.Printf("4:%d\n",a.v);  //2
//}
//
//func main() {
//    aobj := abc{} // new(abc);
//	fmt.Println(aobj)
//    aobj.aaaa();
//    aobj.bbbb();
//    aobj.cccc();
//}
