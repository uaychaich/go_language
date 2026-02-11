package main

import (
	"fmt"
)

var a1 = 5
var a2 int
var a3 int = 5

// a99:=5

func main() {
	var a4 = 5
	var a5 int
	var a6 int = 5
	a7 := 5
	fmt.Println(a1, a2, a3, a4, a5, a6, a7)
	///////////////////////////////////////////////////////
	var b1 bool = true
	var (
		i1 int = 5
		// i2 int8  = 5
		// i3 int16 = 5
		// i4 int32 = 5
		// i5 int64 = 5
	)
	var (
		u1 uint = 5
		// u2 uint8  = 5
		// u3 uint16 = 5
		// u4 uint32 = 5
		// u5 uint64 = 5
	)
	var (
		f1 float32 = 5.2
		// f2 float64 = 5.2
	)
	var (
		c1 complex64 = 5 + 6i
		// c2 complex128 = complex(5, 6)
	)
	var s1 string = "Uaychai"

	const cc1 int = 5
	// cc1 = 7

	fmt.Printf("%t %d %d %9.2f %g %s %d\n", b1, i1, u1, f1, c1, s1, cc1)
	///////////////////////////////////////////////////////
	var arr1 [3]int
	arr1[0] = 5
	arr1[1] = 60
	arr1[2] = 45
	arr2 := [3]int{70, 80, 90}
	arr3 := [...]int{55, 66, 77}
	var arr4 [3][2]int
	arr4[0][0] = 888
	arr5 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("%d %d %d\n", arr2[0], arr3[0], arr5[1][0])
	fmt.Printf("%d %d\n", len(arr1), cap(arr1))
	///////////////////////////////////////////////////////
	var add1 int = 99
	var add2 *int = &add1
	fmt.Println(add1, &add1, add2, *add2)
	add1 = 88
	fmt.Println(add1, &add1, add2, *add2)
	*add2 = 77
	fmt.Println(add1, &add1, add2, *add2)
}
