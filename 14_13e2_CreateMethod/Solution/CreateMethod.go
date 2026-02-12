package main

import "fmt"

func main() {
	var x1 twonums
	x1.x = 5
	x1.y = 2
	fmt.Printf("x1plus=%v | x1sub=%v\n", x1.plus(), x1.sub())

	var x2 twonums
	x2.x = 50
	x2.y = 20
	fmt.Printf("x2plus=%v | x2sub=%v\n", x2.plus(), x2.sub())
}
