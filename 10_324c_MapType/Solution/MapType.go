package main

import "fmt"

func main() {
	var a1 map[int]int
	a1 = make(map[int]int, 100)
	a1[10] = 100
	a1[20] = 200
	for x := range a1 {
		fmt.Printf("a1[%v]=%v\n", x, a1[x])
	}
	fmt.Printf("//////////////////////////////////////\n")
	///////////////////////////////////////////////////////////
	a2 := map[string]int{"uo": 10, "ba": 20, "ap": 30}
	for x := range a2 {
		fmt.Printf("a2[%v]=%v\n", x, a2[x])
	}
	fmt.Printf("//////////////////////////////////////\n")
	///////////////////////////////////////////////////////////
	a1[30] = 300
	a2["zd"] = 40
	for x := range a1 {
		fmt.Printf("a1[%v]=%v\n", x, a1[x])
	}
	for x := range a2 {
		fmt.Printf("a2[%v]=%v\n", x, a2[x])
	}
	fmt.Printf("//////////////////////////////////////\n")
	///////////////////////////////////////////////////////////
	delete(a1, 30)
	delete(a2, "zd")
	for x := range a1 {
		fmt.Printf("a1[%v]=%v\n", x, a1[x])
	}
	for x := range a2 {
		fmt.Printf("a2[%v]=%v\n", x, a2[x])
	}
	fmt.Printf("//////////////////////////////////////\n")
	///////////////////////////////////////////////////////////
}
