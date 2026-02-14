package main

import (
	Package1 "OOP/Package1"
	package1 "OOP/Package1"
	Package2 "OOP/Package2"
	package2 "OOP/Package2"
	"fmt"
	"time"
)

func main() {
	var p *Package1.People = package1.NewPeople("Uay", "Chot", time.Date(2000, 7, 11, 0, 0, 0, 0, time.Local))

	var c *Package1.Customer = package1.NewCustomer("Luke", "Naja", time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local), 3)

	var e *Package2.Employee = package2.NewEmployee("Vader", "Naja", time.Date(2010, 10, 10, 0, 0, 0, 0, time.Local), 35000)

	GetAllinStringFunc(p)
	GetAllinStringFunc(c)
	GetAllinStringFunc(e)
}

func GetAllinStringFunc(i GetAllinStringThing) {
	fmt.Printf("%v\n", i.GetAllinString())
}

type GetAllinStringThing interface {
	GetAllinString() string
}
