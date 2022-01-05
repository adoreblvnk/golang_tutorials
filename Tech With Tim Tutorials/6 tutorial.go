package main

import (
	"fmt"
)

func main() {
	var num1 float64 = 8.5
	var num2 int = 4
	answer := num1 / float64(num2)
	fmt.Printf("%f", answer)
}
