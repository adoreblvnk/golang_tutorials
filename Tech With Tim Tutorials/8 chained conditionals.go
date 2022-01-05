package main

import (
	"fmt"
)

func main() {
	val := (true || false) && !false
	fmt.Printf("%t", val)
}
