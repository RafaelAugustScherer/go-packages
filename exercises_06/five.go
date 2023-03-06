package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Round(
		math.Pi*c.radius*2*100,
	) / 100
}

func (s square) area() float64 {
	return s.side * s.side
}

func info(s shape) {
	fmt.Println("Shape area is:", s.area())
}

func main() {
	c1 := circle{
		radius: 2,
	}

	s1 := square{
		side: 15,
	}

	info(c1)
	info(s1)
}
