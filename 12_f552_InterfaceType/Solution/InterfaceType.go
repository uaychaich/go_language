package main

import "fmt"

func main() {
	i1 := 5
	i2 := "Uaychai"
	i3 := [3]int{1, 2, 3}
	i4 := struct {
		firstname string
		lastname  string
	}{firstname: "Uaychai", lastname: "Chot"}
	i5 := func(x int, y int) int {
		return x + y
	}
	////////////////////////////////////////////////////////////
	var i21 interface{}
	i21 = i1
	fmt.Printf("i21 for i1: %v\n", i21)
	i21 = i2
	fmt.Printf("i21 for i2: %v\n", i21)
	i21 = i3
	fmt.Printf("i21 for i3: %v\n", i21)
	i21 = i4
	fmt.Printf("i21 for i4: %v\n", i21)
	i21 = i5
	fmt.Printf("i21 for i5: %v\n", i21)
	fmt.Printf("///////////////////////////////////////\n")
	////////////////////////////////////////////////////////////
}
