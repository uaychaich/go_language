package main

import (
	"fmt"
)

func main() {
	var i1 int = 5
	if i1 > 5 {
		fmt.Printf("More than 5 | i1:%d\n", i1)
	} else if i1 == 7 {
		fmt.Printf("Equal 7 | i1:%d\n", i1)
	} else {
		fmt.Printf("Less than 7 | i1:%d\n", i1)
	}
	fmt.Printf("i1:%d\n", i1)
	////////////////////////////////////////////////////////////
	if i2 := 5; i2 > 5 {
		fmt.Printf("More than 5 | i2:%d\n", i2)
	} else if i2 == 7 {
		fmt.Printf("Equal 7 | i2:%d\n", i2)
	} else {
		fmt.Printf("Less than 7 | i2:%d\n", i2)
	}
	//fmt.Printf("i2:%d\n",i2)
	////////////////////////////////////////////////////////////
	var i3 int = 5
	switch i3 {
	case 1, 3, 5, 7, 9:
		fmt.Printf("Odd | i3:%d\n", i3)
	case 2, 4, 6, 8, 10:
		fmt.Printf("Even | i3:%d\n", i3)
	default:
		fmt.Printf("Out of range | i3:%d\n", i3)
	}
	////////////////////////////////////////////////////////////
	var i4 int = 7
	switch {
	case i4 > 5:
		fmt.Printf("Mor than 5 | i4:%d\n", i4)
	case i4 < 5:
		fmt.Printf("Less than 5 | i4:%d\n", i4)
	default:
		fmt.Printf("Equal 5 | i4:%d\n", i4)
	}
	////////////////////////////////////////////////////////////
	switch i5 := 7; {
	case i5 > 5:
		fmt.Printf("Mor than 5 | i5:%d\n", i5)
	case i5 < 5:
		fmt.Printf("Less than 5 | i5:%d\n", i5)
	default:
		fmt.Printf("Equal 5 | i5:%d\n", i5)
	}
	//fmt.Printf("i5:%d\n", i5)
	////////////////////////////////////////////////////////////
	var i6 interface{} = 5
	switch i6.(type) {
	case string:
		fmt.Printf("String type | i6=%v\n", i6)
	case int:
		fmt.Printf("Int type | i6=%v\n", i6)
	default:
		fmt.Printf("Not string or int type | | i6=%v\n", i6)
	}
	////////////////////////////////////////////////////////////
	switch i7 := 8; {
	case i7 > 5:
		if i7 == 7 {
			fmt.Printf("Breack Naja")
			break
		}
		fmt.Printf("Mor than 5 | i7:%d\n", i7)
	case i7 < 5:
		fmt.Printf("Less than 5 | i7:%d\n", i7)
	default:
		fmt.Printf("Equal 5 | i7:%d\n", i7)
	}
	//fmt.Printf("i7:%d\n", i7)
}
