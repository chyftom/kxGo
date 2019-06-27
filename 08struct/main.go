//1.结构体类型的定义
// type Circle struct {
// 	x      int
// 	y      int
// 	Radius int
// }
//2.结构体变量的创建
// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// type Circle struct {
// 	x      int
// 	y      int
// 	Radius int
// }

// func main() {
// 	var c Circle = Circle{
// 		x:      100,
// 		y:      100,
// 		Radius: 50, // 注意这里的逗号不能少
// 	}
// 	fmt.Printf("%+v\n", c)
// }
// func main() {
// 	var c1 Circle = Circle{
// 		Radius: 50,
// 	}
// 	var c2 Circle = Circle{}
// 	fmt.Printf("%+v\n", c1)
// 	fmt.Printf("%+v\n", c2)
// }
// func main() {
// 	var c *Circle = &Circle{100, 100, 50}
// 	fmt.Printf("%+v\n", c)
// }
// func main() {
// 	var c *Circle = new(Circle)
// 	fmt.Printf("%+v\n", c)

// }

//最后我们再将三种零值初始化形式放到一起对比观察一下
// var c1 Circle = Circle{}
// var c2 Circle
// var c3 *Circle = new(Circle)

//3.零值结构体和 nil 结构体
// func main() {
// 	var c *Circle = nil
// 	// fmt.Printf("%+v\n", len(c))
// 	fmt.Println(unsafe.Sizeof(c))
// }

//4.结构体的内存大小
// func main() {
// 	var c Circle = Circle{} //= Circle{Radius: 50}
// 	fmt.Println(unsafe.Sizeof(c))
// }

//5.结构体的拷贝
// func main() {
// 	var c Circle = Circle{} //= Circle{Radius: 50}
// 	fmt.Println(unsafe.Sizeof(c))

// 	var c1 Circle = Circle{Radius: 50}
// 	var c2 Circle = c1
// 	fmt.Printf("%+v\n", c1)
// 	fmt.Printf("%+v\n", c2)
// 	c1.Radius = 100
// 	fmt.Printf("%+v\n", c1)
// 	fmt.Printf("%+v\n", c2)

// 	var c3 *Circle = &Circle{Radius: 50}
// 	var c4 *Circle = c3
// 	fmt.Printf("%+v\n", c3)
// 	fmt.Printf("%+v\n", c4)
// 	c3.Radius = 100
// 	fmt.Printf("%+v\n", c3)
// 	fmt.Printf("%+v\n", c4)

// 	fmt.Println(c1 == c2)
// 	fmt.Println(c3 == c4)
// }

//6.无处不在的结构体
// 切片头的结构体形式如下，它在 64 位机器上将会占用 24 个字节

// type slice struct {
//   array unsafe.Pointer  // 底层数组的地址
//   len int // 长度
//   cap int // 容量
// }

// 字符串头的结构体形式，它在 64 位机器上将会占用 16 个字节

// type string struct {
//   array unsafe.Pointer // 底层数组的地址
//   len int
// }

// 字典头的结构体形式

// type hmap struct {
//   count int
//   ...
//   buckets unsafe.Pointer  // hash桶地址
//   ...
// }
//7.结构体中的数组和切片
// type ArrayStruct struct {
// 	value [11]int
// }

// type SliceStruct struct {
// 	value []int
// }

// func main() {
// 	// var as = ArrayStruct{[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1}}
// 	// var ss = SliceStruct{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11}}
// 	fmt.Println(unsafe.Sizeof(1), unsafe.Sizeof(1))
// }
// package main

// import "fmt"
// import "unsafe"

// type ArrayStruct struct {
// 	value [11]int
// }

// type SliceStruct struct {
// 	value []int
// }

// func main() {
// 	var as = ArrayStruct{[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1}}
// 	var ss = SliceStruct{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
// 	fmt.Println(unsafe.Sizeof(as), unsafe.Sizeof(ss))
// }

//8.结构体的参数传递
// package main

// import "fmt"

// type Circle struct {
// 	x      int
// 	y      int
// 	Radius int
// }

// func expandByValue(c Circle) {
// 	c.Radius *= 2
// }

// func expandByPointer(c *Circle) {
// 	c.Radius *= 2
// }

// func main() {
// 	var c = Circle{Radius: 50}
// 	expandByValue(c)
// 	fmt.Println(c)
// 	expandByPointer(&c)
// 	fmt.Println(c)
// }

//9.结构体方法
//10.
//11
