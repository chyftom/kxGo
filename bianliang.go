package main

import "fmt"

const globali int = 24

func doSomething()  {
	fmt.Println("do something")
}

func main() {
	//var s  = 42
	//fmt.Println(s)
	//
	//s2 := 420
	//fmt.Println(s2)

	//for i:=0; i<10; i++ {
	//	doSomething()
	//}

	//var i int
	//fmt.Println(i)
	//
	//var boolValue bool
	//fmt.Println(boolValue)

	//const locali int = 42
	//fmt.Println(globali, locali)
	//var value int = 42
	//var pointer *int = &value
	//fmt.Println(pointer, *pointer)

	//var value int = 42
	//var p1 *int = &value
	//var p2 **int = &p1
	//var p3 ***int = &p2
	//fmt.Println(p1, p2, p3)
	//fmt.Println(*p1, **p2, ***p3)

	// ***int (0xc00000a0c0) --> **int (0xc000006028) --> *int (0xc000006030)--> int (42)

	// 有符号整数，可以表示正负
	//var a int8 = 1 // 1 字节
	//var b int16 = 2 // 2 字节
	//var c int32 = 3 // 4 字节
	//var d int64 = 4 // 8 字节
	//fmt.Println(a, b, c, d)

	// 无符号整数，只能表示非负数
	//var ua uint8 = 1
	//var ub uint16 = 2
	//var uc uint32 = 3
	//var ud uint64 = 4
	//fmt.Println(ua, ub, uc, ud)

	//var e int = 5
	//var ue uint = 5
	//fmt.Println(e, ue)
	//// bool 类型
	//var f bool = true
	//fmt.Println(f)
	//
	//// 字节类型
	//var j [3]byte = [3]byte{'a','b','c'}
	//
	//fmt.Println(j)
	//
	//// 字符串类型
	//var g string = "abcdefg"
	//fmt.Println(g)
	//
	//// 浮点数
	//var h float32 = 3.14
	//var i float64 = 3.141592653
	//fmt.Println(h, i)

	var j [3]byte = [3]byte{'a','b','c'}
	for i := 0;i< len(j);i++{
		fmt.Println(j[i])
	}
	fmt.Println(j)


}
