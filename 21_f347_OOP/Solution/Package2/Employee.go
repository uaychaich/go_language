package package2

import (
	Package1 "OOP/Package1"
	"fmt"
	"time"
)

type Employee struct {
	Package1.People
	salary uint32
}

func (e *Employee) SetSalary(salary uint32) {
	if salary > 0 {
		e.salary = salary
	} else {
		panic("Salary must be more than 0")
	}
}

func (e *Employee) GetSalary() uint32 {
	return e.salary
}

func (e *Employee) GetAllinString() string {
	return fmt.Sprintf("Employee\nFullName: %v\nBirthday: %v\nSalary: %v\n\n",
		e.FirstName+" "+e.LastName, e.GetBirthday().String(), e.GetSalary())
}

func NewEmployee(firstname string, lastname string, birthday time.Time, salary uint32) *Employee {
	e := Employee{}
	e.FirstName = firstname
	e.LastName = lastname
	e.SetBirthday(birthday)
	e.SetSalary(salary)
	return &e
}
