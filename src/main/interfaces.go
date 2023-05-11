package main

import "fmt"

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	base   float64
	height float64
	color  string
}

func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}

func (t Triangle) Perimeter() float64 {
	return 0
}

func main() {
	printMe("Test")
}

func printMe(val interface{}) {
	fmt.Println(val)
}
