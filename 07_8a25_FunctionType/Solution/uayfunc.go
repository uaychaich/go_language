package main

import (
	"fmt"
)

func UayFunc1() {
	fmt.Printf("Hello World!!\n")
}

func UayFunc2(x int, y int) {
	fmt.Printf("x=%v | y=%v\n", x, y)
}

func UayFunc3(x int, y int) int {
	return x + y
}

func UayFunc4(x int, y int) (int, int) {
	return x + y, x - y
}

func UayFunc5(x int, y int) (a int, b int) {
	a = x + y
	b = x - y
	return
}

func UayFunc6(x int, y int, z ...int) (a int, b int, c int) {
	a = x + y
	b = x - y
	for item := range z {
		c = c + z[item]
	}
	return
}

func UayFunc7(x int, y int, z func(int, int) int) int {
	return z(x, y)
}

func UayFunc8(x int) (somefunc func(int, int) int) {
	if x == 1 {
		somefunc = func(a int, b int) int { return a + b }
	} else {
		somefunc = func(a int, b int) int { return a - b }
	}
	return
}

func UayFunc9(x int, y int) (a int) {
	defer func() {
		a += 7
	}()
	a = x + y
	return
}
