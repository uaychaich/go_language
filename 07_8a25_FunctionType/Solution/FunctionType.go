package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!!")
	UayFunc1()
	UayFunc2(5, 2)

	a3 := UayFunc3(5, 2)
	fmt.Printf("a3=%v\n", a3)

	a41, a42 := UayFunc4(5, 2)
	fmt.Printf("a41=%v | a42=%v\n", a41, a42)

	a51, a52 := UayFunc5(5, 2)
	fmt.Printf("a51=%v | a52=%v\n", a51, a52)

	a61, a62, a63 := UayFunc6(5, 2, 3, 4, 5)
	fmt.Printf("a61=%v | a62=%v | a63=%v\n", a61, a62, a63)

	a71 := UayFunc7(5, 2, func(a int, b int) int { return a + b })
	a72 := UayFunc7(5, 2, func(a int, b int) int { return a - b })
	somefunc := func(a int, b int) int { return a * b }
	a73 := UayFunc7(5, 2, somefunc)
	fmt.Printf("a71=%v | a72=%v | a73=%v\n", a71, a72, a73)

	plusfunc := UayFunc8(1)
	subfunc := UayFunc8(2)
	a81 := plusfunc(5, 2)
	a82 := subfunc(5, 2)
	fmt.Printf("a81=%v | a82=%v\n", a81, a82)

	a91 := UayFunc9(5, 2)
	fmt.Printf("a91=%v\n", a91)

	a101 := UayFunc10(5, 2)
	fmt.Printf("a101=%v\n", a101)
	fmt.Printf("/////////////////////////////////////////////\n")
	///////////////////////////////////////////////////////////////////////////////
	var a111 func(int, int) int
	a111 = func(x int, y int) int { return x + y }
	fmt.Printf("a111=%v\n", a111(5, 2))
	fmt.Printf("/////////////////////////////////////////////\n")
	a112 := func(x int, y int) int { return x + y }
	fmt.Printf("a112=%v\n", a112(5, 2))
	fmt.Printf("/////////////////////////////////////////////\n")
}
