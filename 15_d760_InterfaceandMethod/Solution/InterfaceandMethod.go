package main

import "fmt"

func main() {
	var a1 circle
	a1.radius = 20

	var a2 triangle
	a2.base = 15
	a2.height = 20

	showarea1(a1)
	showarea2(a2)
	showarea3(a1)
	showarea3(a2)
}

func showarea1(c circle) {
	fmt.Printf("Area is %v\n", c.area())
}

func showarea2(t triangle) {
	fmt.Printf("Area is %v\n", t.area())
}

func showarea3(z areathing) {
	fmt.Printf("Area is %v\n", z.area())
}
