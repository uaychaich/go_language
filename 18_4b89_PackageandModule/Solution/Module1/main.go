package main

import (
	"fmt"

	package1 "uay.com/Module1/Package1"
	package2 "uay.com/Module1/Package2"
)

func main() {
	fmt.Printf("Hello Worlds\n")
	var i1 package1.Animal = package1.Animal{Weight: 53, Height: 100}
	var i2 package2.People = package2.People{Firstname: "Uay", Lastname: "Chot"}
	fmt.Printf("i1=%v\n", i1)
	fmt.Printf("i2=%v\n", i2)
}
