package main

import (
	"fmt"
	"math"
)

func main() {

	rect := Rectangle{Width: 10, Height: 10}
	rect.Area()
	rect.Perimeter()
	cir := Circle{Radius: 10}
	cir.Area()
	cir.Perimeter()
}

type Shape interface {
	Area()
	Perimeter()
}
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	rst := r.Width * r.Height
	fmt.Println("Area of Rectangle : ", rst)
	return rst
}
func (r Rectangle) Perimeter() float64 {
	rst := 2 * (r.Width + r.Height)
	fmt.Println("Perimeter of Rectangle : ", rst)
	return rst
}
func (r Circle) Area() float64 {
	rst := math.Pi * r.Radius * r.Radius
	fmt.Println("Area of Circle : ", rst)
	return rst
}
func (r Circle) Perimeter() float64 {
	rst := 2 * math.Pi * r.Radius
	fmt.Println("Perimeter of Circle : ", rst)
	return rst
}
