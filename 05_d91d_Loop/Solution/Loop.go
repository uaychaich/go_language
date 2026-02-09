package main

import (
	"fmt"
)

func main() {
	var i1 int = 0
	for i1 < 5 {
		i1 = i1 + 1
		fmt.Printf("i1=%d\n", i1)
	}
	fmt.Printf("---\n")
	/////////////////////////////////////////////////
	for i2 := 10; i2 <= 50; i2 = i2 + 10 {
		fmt.Printf("i2=%d\n", i2)
	}
	fmt.Printf("---\n")
	/////////////////////////////////////////////////
	i3 := []int{100, 200, 300, 400, 500}
	for j := range i3 {
		fmt.Printf("i3[%d]=%d\n", j, i3[j])
	}
	fmt.Printf("---\n")
	/////////////////////////////////////////////////
	i4 := []int{100, 200, 300, 400, 500}
	for j := range i4 {
		if j == 2 {
			continue
		}
		if j == 3 {
			break
		}
		fmt.Printf("i4[%d]=%d\n", j, i4[j])
	}
	fmt.Printf("---\n")
	/////////////////////////////////////////////////
}
