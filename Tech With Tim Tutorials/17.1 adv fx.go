package main

import (
	"fmt"
)

func returnFunc(x string) func() {
	// function closure
	return func() { fmt.Println(x) }
}

func main() {
	returnFunc("hello")()
	x := returnFunc("goodbye")
	x()
}
