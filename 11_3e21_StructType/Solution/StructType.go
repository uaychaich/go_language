package main

import (
	"fmt"
)

func main() {
	var st1 struct {
		x   int
		y   int
		cal func(int, int) int
	}
	st1.x = 25
	st1.y = 80
	st1.cal = func(i1 int, i2 int) int { return i1 + i2 }
	fmt.Printf("st1.x=%d | st1.y=%d | cal=%d\n", st1.x, st1.y, st1.cal(st1.x, st1.y))
	fmt.Printf("/////////////////////////////////////////////\n")
	////////////////////////////////////////////////////////////////////////////
	st2 := struct {
		x   int
		y   int
		cal func(int, int) int
	}{x: 25, y: 80, cal: func(i1 int, i2 int) int { return i1 + i2 }}
	fmt.Printf("st2.x=%d | st2.y=%d | cal=%d\n", st2.x, st2.y, st1.cal(st1.x, st1.y))
	fmt.Printf("/////////////////////////////////////////////\n")
	////////////////////////////////////////////////////////////////////////////
}
