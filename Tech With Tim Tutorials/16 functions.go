package main

import "fmt"

func calc(x, y int) (add, subtract int) {
	defer fmt.Println("hello")
	add = x + y
	subtract = x - y
	fmt.Println("before return")
	return
}

func main() {
	ans1, ans2 := calc(6, 7)
	fmt.Println(ans1, ans2)
}
