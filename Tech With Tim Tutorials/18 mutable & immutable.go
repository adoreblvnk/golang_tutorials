package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 1000
}

func main() {
	// // slice & map is mutable
	// var x map[string]int = map[string]int{"hello": 3}
	// y := x
	// y["y"] = 100
	// fmt.Println(x, y)

	// more mutable
	var x []int = []int{3, 4, 5}
	fmt.Println(x)
	changeFirst(x)
	fmt.Println(x)
}
