package interfaces

import (
	"fmt"
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

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.R, 2.0)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func (s *Square) Area() float64 {
	return math.Pow(s.S, 2.0)
}

func GetPrice(land GeometricShapeArea) float64 {
	return 10 * (land).Area()
}

func isGeometricShape(s GeometricShape) bool {
	return true
}

func InterfaceTest() {
	// use a *T to get price
	var sq = Square{S: 4}
	var sqPrice = GetPrice(&sq)
	fmt.Println(sqPrice)

	// use a T or *T to get price
	var circ = Circle{R: 10}
	var sqCirc = GetPrice(circ)
	fmt.Println(sqCirc)
	circ = Circle{R: 100}
	sqCirc = GetPrice(&circ)
	fmt.Println(sqCirc)

	// use an interface to get price
	var geoArea GeometricShapeArea = &Rectangle{2, 4}
	var rectPrice = GetPrice(geoArea)
	fmt.Println(rectPrice)

	// isGeometricShape is available only to those types that implement both GeometricShapeArea and GeometricShapePerimeter
	isGeometricShape(&circ)
}
