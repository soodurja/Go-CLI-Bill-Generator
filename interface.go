package main

import (
	"fmt"
	"math"
)

// Shape interface
type shape interface {
	getArea() float64
	getCircumference() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// Sqaure methods
func (s square) getArea() float64 {
	return s.length * s.length
}

func (s square) getCircumference() float64 {
	return 4 * s.length
}

// Circle methods
func (c circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) getCircumference() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("The area of %T is %0.2f\n", s, s.getArea())
	fmt.Printf("The circumference of %T is %0.2f\n", s, s.getCircumference())
}

// func main() {
// 	shapes := []shape{
// 		square{length: 15},
// 		circle{radius: 5},
// 		square{length: 10},
// 		circle{radius: 11},
// 	}

// 	for _, v := range shapes {
// 		printShapeInfo(v)
// 	}
// }
