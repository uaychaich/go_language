package main

import "fmt"

func main() {
	var i1 uayint = 5
	var i2 uaystring = "Uaychai"
	var i3 uayarray = uayarray{10, 20, 30}
	var i4 uayslice = uayslice{100, 200, 300}
	var i5 uaymap = uaymap{"uo": 10, "ba": 20, "ap": 30}
	var i6 uayfunc = func(i1, i2 int) int { return i1 + i2 }
	var i7 uaystruct = uaystruct{firstname: "Uaychai", lastname: "Chotja"}

	fmt.Printf("i1=%v\n", i1)
	fmt.Printf("i2=%v\n", i2)
	fmt.Printf("i3=%v\n", i3)
	fmt.Printf("i4=%v\n", i4)
	fmt.Printf("i5=%v\n", i5)
	fmt.Printf("i6=%v\n", i6)
	fmt.Printf("i7=%v\n", i7)
}
