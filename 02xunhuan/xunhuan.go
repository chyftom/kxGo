package main

import "fmt"

var i int

func main() {
	for {
		i = i + 1
		fmt.Println("hello world!", i)
	}

	// ==
	for true {
		fmt.Println("hello world!")
	}
}

// package main

// import "fmt"

// func main() {
// 	fmt.Println(prize1(60))
// 	fmt.Println(prize2(60))
// }

// //switch 有两种匹配模式，一种是变量值匹配，一种是表达式匹配。

// // 值匹配
// func prize1(score int) string {
// 	switch score / 10 {
// 	case 0, 1, 2, 3, 4, 5:
// 		return "差"
// 	case 6, 7:
// 		return "及格"
// 	case 8:
// 		return "良"
// 	default:
// 		return "优"
// 	}
// }

// // 表达式匹配
// func prize2(score int) string {
// 	// 注意 switch 后面什么也没有
// 	switch {
// 	case score < 60:
// 		return "差"
// 	case score < 80:
// 		return "及格"
// 	case score < 90:
// 		return "良"
// 	default:
// 		return "优"
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println(sign(max(min(24, 42), max(24, 42))))
// }

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a int, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func sign(a int) int {
// 	if a > 0 {
// 		return 1
// 	} else if a < 0 {
// 		return -1
// 	} else {
// 		return 0
// 	}
// }
