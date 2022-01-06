package main

import "fmt"

func main() {
	// first style
	// x := 0
	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// second style
	// for x := 0; x < 5; x++ {
	// 	fmt.Println(x)
	// }

	// 3rd style
	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			continue // break
		}
		// fmt.Println("N")
	}
}
