package main

import "fmt"

func main() {
	// // slices
	// var x [5]int = [5]int{1, 2, 3, 4, 5}
	// var s []int = x[1:3]
	// fmt.Println(cap(s)) // capacity is from values 2 to 5
	// fmt.Println(s[1:])  // reslicing

	// // slicing easier
	// var a []int = []int{5, 6, 7, 8, 9}
	// a = append(a, 10)
	// fmt.Println(a)

	// using make() to create slice
	a := make([]int, 5)
	fmt.Println(a)
}
