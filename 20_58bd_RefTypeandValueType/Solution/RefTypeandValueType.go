package main

import (
	"fmt"
)

func main() {
	var a1 int = 5
	var a2 float64 = 5.2
	var a3 bool = true
	var a4 string = "Uaychai"
	var a5 [3]int = [3]int{1, 2, 3}
	var a6 []int = []int{1, 2, 3, 4}
	var a7 map[string]int = map[string]int{"yoda": 53, "luke": 43}
	a8 := struct {
		x int
		y int
	}{x: 7, y: 8}
	a9 := struct {
		x int
		y int
		z []int
	}{x: 7, y: 8, z: []int{1, 2, 3, 4}}
	var a10 chan int = make(chan int, 2)
	a10 <- 7
	var a11 func(int, int) int = func(x int, y int) int { return x + y }
	///////////////////////////////////////////////////////////////////////
	b1 := a1
	fmt.Printf("before | a1=%v | b1=%v : ", a1, b1)
	b1 = 7
	fmt.Printf("after | a1=%v | b1=%v\n", a1, b1)
	///////////////////////////////////////////////////////////////////////
	b2 := a2
	fmt.Printf("before | a2=%v | b2=%v : ", a2, b2)
	b2 = 7.5
	fmt.Printf("after | a2=%v | b2=%v\n", a2, b2)
	///////////////////////////////////////////////////////////////////////
	b3 := a3
	fmt.Printf("before | a3=%v | b3=%v : ", a3, b3)
	b3 = false
	fmt.Printf("after | a3=%v | b3=%v\n", a3, b3)
	///////////////////////////////////////////////////////////////////////
	b4 := a4
	fmt.Printf("before | a4=%v | b4=%v : ", a4, b4)
	b4 = "Yoda"
	fmt.Printf("after | a4=%v | b4=%v\n", a4, b4)
	///////////////////////////////////////////////////////////////////////
	b5 := a5
	fmt.Printf("before | a5=%v | b5=%v : ", a5, b5)
	b5[1] = 700
	fmt.Printf("after | a5=%v | b5=%v\n", a5, b5)
	///////////////////////////////////////////////////////////////////////
	b6 := a6
	fmt.Printf("before | a6=%v | b6=%v : ", a6, b6)
	b6[1] = 700
	fmt.Printf("after | a6=%v | b6=%v\n", a6, b6)
	///////////////////////////////////////////////////////////////////////
	b7 := a7
	fmt.Printf("before | a7=%v | b7=%v : ", a7, b7)
	b7["yoda"] = 700
	fmt.Printf("after | a7=%v | b7=%v\n", a7, b7)
	///////////////////////////////////////////////////////////////////////
	b8 := a8
	fmt.Printf("before | a8=%v | b8=%v : ", a8, b8)
	b8.x = 700
	fmt.Printf("after | a8=%v | b8=%v\n", a8, b8)
	///////////////////////////////////////////////////////////////////////
	b9 := a9
	fmt.Printf("before | a9=%v | b9=%v : ", a9, b9)
	b9.x = 700
	b9.z[1] = 800
	fmt.Printf("after | a9=%v | b9=%v\n", a9, b9)
	///////////////////////////////////////////////////////////////////////
	b10 := a10
	fmt.Printf("before | len(a10)=%v | len(b10)=%v : ", len(a10), len(b10))
	b10 <- 9
	fmt.Printf("after | len(a10)=%v | len(b10)=%v\n", len(a10), len(b10))
	///////////////////////////////////////////////////////////////////////
	b11 := a11
	fmt.Printf("before | a11=%v | b11=%v : ", a11, b11)
	b11 = func(x int, y int) int { return x - y }
	fmt.Printf("after | a11=%v | b11=%v\n", a11, b11)
	///////////////////////////////////////////////////////////////////////
}
