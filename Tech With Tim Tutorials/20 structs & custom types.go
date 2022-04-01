package main

import "fmt"

type Point struct {
	x int32
	y int32
}

func changeX(pt *Point) {
	pt.x = 100
}

type Circle struct {
	radius float64
	*Point
}

func main() {
	// // basic struct declaration
	// var p1 Point = Point{1, 2}
	// var p2 Point = Point{-5, 2}
	// fmt.Println(p1.x, p2.x)
	// fmt.Println(p2.y)

	// p1 := &Point{y: 3}
	// fmt.Println(p1)
	// changeX(p1) // NOTE: for strcuts, you do not need to dereference it
	// fmt.Println(p1)

	// NOTE: with a nested struct: it's important to reference the pointer, not just the struct
	c1 := Circle{4.43, &Point{4, 5}}
	fmt.Println(c1.x)
	fmt.Println(c1.radius)
}
