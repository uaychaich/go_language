package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var a1 float64 = 5
	var a2 float64 = 2
	var a3 int = 5
	var a4 int = 2
	fmt.Printf("%f %f %f %f %d\n", a1+a2, a1-a2, a1*a2, a1/a2, a3%a4)
	//////////////////////////////////////////////////////////////////
	var b1 float64 = 5
	var b2 float64 = 2
	fmt.Printf("%t %t %t %t %t %t\n", b1 == b2, b1 != b2, b1 < b2, b1 <= b2, b1 > b2, b1 >= b2)
	//////////////////////////////////////////////////////////////////
	var c1 bool = true
	var c2 bool = false
	fmt.Printf("%t %t %t\n", c1 && c2, c1 || c2, !c2)
	//////////////////////////////////////////////////////////////////
	z1 := []int{5, 7, 9}
	z1 = append(z1, 8)
	z2 := make([]int, len(z1), 100)
	copy(z2, z1)
	println(z1[0], z1[1], z1[2], z2[0], z2[1], z2[2])
	//////////////////////////////////////////////////////////////////
	j1 := 5
	j2 := 5.2
	j3 := "52"
	j4 := strconv.Itoa(j1)
	j5, j6 := strconv.Atoi(j3)
	j7, j8 := strconv.ParseInt(j3, 10, 64)
	j9 := float32(j1)
	j10 := int(j2)
	fmt.Printf("%T %T %T %T %T\n", j4, j5, j7, j9, j10)
	fmt.Println(reflect.TypeOf(j9), reflect.ValueOf(j9).Kind(), j6, j8)
	//////////////////////////////////////////////////////////////////
}
