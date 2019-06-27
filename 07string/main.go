package main

import "fmt"

// func main() {
// 	var s = "嘻哈china"
// 	fmt.Println(len(s))
// 	for i := 0; i < len(s); i++ {
// 		fmt.Println(s[i])
// 	}

// }

// func main() {
// 	// var s = "嘻哈china"
// 	// for codepoint, runeValue := range s {
// 	// 	fmt.Printf("%d %d ", codepoint, int32(runeValue))
// 	// }
// 	var s1 = "hello" // 静态字面量
// 	var s2 = ""
// 	for i := 0; i < 10; i++ {
// 		s2 += s1 // 动态构造
// 	}
// 	// fmt.Println(len(s1))
// 	// fmt.Println(len(s2))

// 	// for i := 0; i < len(s1); i++ {
// 	// 	fmt.Printf("%x ", s1[i])
// 	// }
// 	// fmt.Println("next-->")
// 	// for i := 0; i < len(s2); i++ {
// 	// 	fmt.Printf("%x ", s2[i])
// 	// }
// 	fmt.Println("next1-->")
// 	fmt.Printf("%x ", &s1)
// 	// fmt.Printf("%x ", &s1[0])
// 	fmt.Println("next2-->")
// 	fmt.Printf("%x ", &s2)
// 	// fmt.Printf("%x ", &s2[0])
// }
func main() {
	// var s1 = "hello world中"
	// var b = []byte(s1) // 字符串转字节切片
	// var s2 = string(b) // 字节切片转字符串
	// fmt.Println(b)
	// fmt.Println(s2)

	// var s1 = "hello world"
	// var s2 = s1[3:8]
	//而字符串的底层字节数组是只读的
	// s1[0] = 'H'
	//cannot assign to s1[0]
	// fmt.Println(s2)

	var s1 = "hello world"

	//字节切片的底层数组内容是可以修改的
	var b = []byte(s1) // 字符串转字节切片
	b[0] = 'H'
	var s2 = string(b)
	fmt.Println(s1, s2)
}
