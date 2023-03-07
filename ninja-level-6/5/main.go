package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return float64(s.length * s.width)
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type) {
	case square:
		fmt.Println(s.(square).area())
	case circle:
		fmt.Println(s.(circle).area())
	}

}
func main() {
	c1 := circle{
		radius: 3,
	}
	s1 := square{
		length: 4,
		width:  6,
	}
	info(c1)
	info(s1)

}
