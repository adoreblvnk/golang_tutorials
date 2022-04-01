package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!!"
}

func main() {
	// // basic deference + pointer
	// x := 7
	// y := &x
	// fmt.Println(x, y)
	// *y = 8
	// fmt.Println(x, y)

	// // dereference function
	// toChange := "hello"
	// fmt.Println(toChange)
	// changeValue(&toChange)
	// fmt.Println(toChange)

	// gets the memory address of var, then gets the value of the memory address
	toChange := "hello"
	var pointer *string = &toChange
	fmt.Print(*pointer)
}
