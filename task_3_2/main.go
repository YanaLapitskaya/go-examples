package main

import (
	"fmt"
	"math"
)

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	a float64
}

func (s Square) area() float64 {
	return s.a * s.a
}

func (s Square) perimeter() float64 {
	return s.a * 4
}

type Circle struct {
	r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func main() {
	var s Figure = Square{5}
	var c Figure = Circle{4}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())
}
