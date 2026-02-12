package main

import "math"

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

type areathing interface {
	area() float64
}
