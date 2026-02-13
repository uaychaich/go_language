package main

import (
	"fmt"

	package1_1 "uay.com/Module1/Package1"
	package1_2 "uay.com/Module1/Package2"
	package2_1 "uay.com/Module2/Package1"
	package2_2 "uay.com/Module2/Package2"
)

func main() {
	fmt.Printf("Hello Worlds\n")
	var i1 package1_1.Animal = package1_1.Animal{Weight: 53, Height: 100}
	var i2 package1_2.People = package1_2.People{Firstname: "Uay", Lastname: "Chot"}
	var i3 package2_1.Circle = package2_1.Circle{Radius: 23.5}
	var i4 package2_2.Triangle = package2_2.Triangle{Base: 30, Height: 25}

	fmt.Printf("i1=%v\n", i1)
	fmt.Printf("i2=%v\n", i2)
	fmt.Printf("i3=%v\n", i3)
	fmt.Printf("i4=%v\n", i4)
}
