package main

import "fmt"

func main() {
	ans := -1

	// switch with values
	// switch ans {
	// case 1, -1:
	// 	fmt.Println("absolute one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("not a case")
	// }

	// switch with expressions
	switch {
	case ans > 0:
		fmt.Println("greater than 0")
	case ans < 0:
		fmt.Println("less than 0")
	default:
		fmt.Println("zero")
	}
}
