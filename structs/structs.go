package interfaces

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	base   float64
	height float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func Perimeter(shape Rectangle) float64 {

	return 2 * (shape.width + shape.height)

}

func Area(shape Rectangle) float64 {
	return shape.width * shape.height
}
