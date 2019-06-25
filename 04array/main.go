package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	for index := range a {
		fmt.Println(index, a[index])
	}
	for index, value := range a {
		fmt.Println(index, value)
	}
}

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
