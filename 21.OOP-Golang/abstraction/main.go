package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	var shape Shape

	shape = Rectangle{width: 5, height: 3}
	fmt.Println("Прямоугольник:")
	fmt.Println("Площадь:", shape.Area())
	fmt.Println("Периметр:", shape.Perimeter())

	shape = Circle{radius: 4}
	fmt.Println("Круг:")
	fmt.Println("Площадь:", shape.Area())
	fmt.Println("Периметр:", shape.Perimeter())
}
