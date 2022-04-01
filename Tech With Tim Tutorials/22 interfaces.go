package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type shape_2 interface {
	perimeter() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	circle_one := circle{4.5}
	rect_one := rect{5, 6}
	shapes := []shape{&circle_one, &rect_one}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

}
