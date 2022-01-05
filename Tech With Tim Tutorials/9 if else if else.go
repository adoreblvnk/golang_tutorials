package main

import (
	"fmt"
)

func main() {
	age := 15
	if age >= 18 {
		fmt.Println("you can ride alone.")
	} else if age >= 14 {
		fmt.Println("you can ride with a parent.")
	} else {
		fmt.Println("you cannot ride.")
	}
}
