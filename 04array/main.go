package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4, 5}
	fmt.Println(s1, len(s1), cap(s1))

	// 对满容的切片进行追加会分离底层数组
	// var s2 = append(s1, 6)
	// fmt.Println(s1, len(s1), cap(s1))
	// fmt.Println(s2, len(s2), cap(s2))

	// 对非满容的切片进行追加会共享底层数组
	// var s3 = append(s2, 7)
	// fmt.Println(s2, len(s2), cap(s2))
	// fmt.Println(s3, len(s3), cap(s3))
}

// func main() {
// 	var a = [5]int{1, 2, 3, 4, 5}

// 	// var s = []int{1,2,3,4,5} 切片没有长度
// 	for index := range a {
// 		fmt.Println(index, a[index])
// 	}
// 	for index, value := range a {
// 		fmt.Println(index, value)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var a = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	var b [9]int
// 	b = a
// 	a[0] = 12345
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// package main

// import "fmt"

// func main() {
// 	var a [9]int
// 	fmt.Println(a)
// }
