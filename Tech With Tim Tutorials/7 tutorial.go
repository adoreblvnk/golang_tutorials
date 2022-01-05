package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 6.4
	val := float64(x)+1.5 != y
	fmt.Printf("%t", val)

	// comparing ascii value of characters
	a := "a"
	b := "b"
	value := a < b
	fmt.Printf("%t", value)
}
