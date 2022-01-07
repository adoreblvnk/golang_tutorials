package main

import "fmt"

func main() {
	// // using the range keyword
	// var a []int = []int{1, 13, 4154, 453, 6}
	// for _, element := range a { // underscore to represent anonymous variable, to use range without using index
	// 	fmt.Printf("%d \n", element)
	// }

	// checking for duplicates
	// NOTE: there is an error if duplicate count is 3 or more
	var a []int = []int{1, 13, 4154, 453, 1, 1, 2, 2}
	for i, element := range a {
		for _, element2 := range a[i+1:] {
			if element2 == element {
				fmt.Println(element)
				break
			}
		}
	}
}
