package main

// import "fmt"

// func main() {
// 	var m map[int]string = make(map[int]string)
// 	fmt.Println(m, len(m))
// }

// func main() {
// 	var m map[int]string = map[int]string{
// 		90: "优秀",
// 		80: "良好",
// 		60: "及格", // 注意这里逗号不可缺少，否则会报语法错误
// 	}
// 	fmt.Println(m, len(m))
// }

// func main() {
// 	var fruits = map[string]int{
// 		"apple":  2,
// 		"banana": 5,
// 		"orange": 8,
// 	}
// 	// 读取元素
// 	var score = fruits["banana"]
// 	fmt.Println(score)

// 	// 增加或修改元素
// 	fruits["pear"] = 3
// 	fmt.Println(fruits)

// 	// 删除元素
// 	delete(fruits, "pear")
// 	fmt.Println(fruits)
// }

// func main() {
// 	var fruits = map[string]int{
// 		"apple":  2,
// 		"banana": 5,
// 		"orange": 8,
// 	}

// 	var score, ok = fruits["durin"]
// 	if ok {
// 		fmt.Println(score)
// 	} else {
// 		fmt.Println("durin not exists")
// 	}

// 	fruits["durin"] = 0
// 	score, ok = fruits["durin"]
// 	if ok {
// 		fmt.Println(score)
// 	} else {
// 		fmt.Println("durin still not exists")
// 	}
// }

//字典的遍历
// func main() {
// 	var fruits = map[string]int{
// 		"apple":  2,
// 		"banana": 5,
// 		"orange": 8,
// 	}

// 	for name, score := range fruits {
// 		fmt.Println(name, score)
// 	}

// 	for name := range fruits {
// 		fmt.Println(name)
// 	}
// }

//可以使用 unsafe 包提供的 Sizeof 函数来计算一个变量的大小

import (
	"fmt"
	// "unsafe"
)

// func main() {
// 	var m = map[string]int{
// 		"apple":  2,
// 		"pear":   3,
// 		"banana": 5,
// 	}
// 	fmt.Println(unsafe.Sizeof(m))
// }

func main() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}

	//这里 len 是数组的长度并且也是切片的初始长度。
	var names = make([]string, 0, len(fruits))
	var scores = make([]int, 0, len(fruits))
	//[apple banana orange] [2 5 8]
	fmt.Println(len(fruits))

	// var names = make([]string, len(fruits))
	// var scores = make([]int, len(fruits))
	//[   banana orange apple] [0 0 0 5 8 2]
	fmt.Println(names, scores)
	for name, score := range fruits {
		fmt.Println(name, score)
		names = append(names, name)
		scores = append(scores, score)
	}

	fmt.Println(names, scores)
}
