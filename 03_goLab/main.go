package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Redius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Redius * c.Redius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Redius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func main() {
	var myShape Shape

	myShape = Circle{2}

	fmt.Println("Area of Circle: ",myShape.Area());
	fmt.Println("Perimeter of Circle: ",myShape.Perimeter());

	myShape=Square{2}
	
	fmt.Println("Area of Square: ",myShape.Area());
	fmt.Println("Perimeter of Square: ",myShape.Perimeter());

}