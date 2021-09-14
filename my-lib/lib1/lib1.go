package lib1

import (
	"math"
)

func Add(a, b int) int {
	return a + b
}

type Rectangle struct {
	L float64
	B float64
}

type Circle struct {
	R float64
}

type Square struct {
	S float64
}

type GeometricShapeArea interface {
	Area() float64
}

type GeometricShapePerimeter interface {
	Perimeter() float64
}

type GeometricShape interface {
	Area() float64
	Perimeter() float64
}

func (r *Rectangle) Area() float64 {
	return r.L * r.B
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.R, 2.0)
}

func (s *Square) Area() float64 {
	return math.Pow(s.S, 2.0)
}

func GetPrice(land GeometricShapeArea) float64 {
	return 10 * land.Area()
}
