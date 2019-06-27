package main

import "fmt"

func main() {
	var s = make([]int, 5, 8)
	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}
	fmt.Println(s)
	var d = make([]int, 3, 6)
	var n = copy(d, s)
	fmt.Println(n, d)
}

// func main() {
// 	var s1 = []int{1, 2, 3, 4, 5}
// 	fmt.Println(s1, len(s1), cap(s1))

// 	// 对满容的切片进行追加会分离底层数组
// 	var s2 = append(s1, 6)
// 	fmt.Println(s1, len(s1), cap(s1))
// 	fmt.Println(s2, len(s2), cap(s2))

// 	// 对非满容的切片进行追加会共享底层数组
// 	var s3 = append(s2, 7)
// 	fmt.Println(s2, len(s2), cap(s2))
// 	fmt.Println(s3, len(s3), cap(s3))
// }

// func main() {
// 	var s1 []int
// 	var s2 []int = []int{}
// 	var s3 []int = make([]int, 0)

// 	if s1 == nil {
// 		fmt.Println("s1 is null")
// 		fmt.Println(len(s1))
// 	}
// 	if s2 == nil {
// 		fmt.Println("s2 is null")
// 	}
// 	if s3 == nil {
// 		fmt.Println("s3 is null")
// 	}
// 	for i := 0; i < len(s1); i++ {

// 	}

// 	// for i := 0; i < len(squares); i++ {
// 	// 	squares[i] = (i + 1) * (i + 1)
// 	// }

// 	for index := range s1 {
// 		fmt.Println(index, s1[index])
// 	}

// 	fmt.Println(s1, s2, s3)
// 	fmt.Println(len(s1), len(s2), len(s3))
// 	fmt.Println(cap(s1), cap(s2), cap(s3))
// }

// func main() {
// 	var s []int = []int{1, 2, 3, 4, 5} // 满容的
// 	fmt.Println(s, len(s), cap(s))
// }

// package main

// import "fmt"

// func main() {
// 	var s1 []int = make([]int, 5, 8)
// 	var s2 []int = make([]int, 8) // 满容切片
// 	fmt.Println(s1)
// 	fmt.Println(s2)
// }
