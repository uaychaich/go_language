package package1

import (
	"fmt"
	"time"
)

type Customer struct {
	People
	star uint8
}

func (c *Customer) SetStar(star uint8) {
	if star >= 1 && star <= 5 {
		c.star = star
	} else {
		panic("star is between 1 and 5")
	}
}

func (c *Customer) GetStar() uint8 {
	return c.star
}

func (c *Customer) GetAllinString() string {
	return fmt.Sprintf("Customer\nFullName: %v\nBirthday: %v\nStar: %v\n\n",
		c.getfullname(), c.GetBirthday().String(), c.GetStar())
}

func NewCustomer(firstname string, lastname string, birthday time.Time, star uint8) *Customer {
	c := Customer{}
	c.FirstName = firstname
	c.LastName = lastname
	c.SetBirthday(birthday)
	c.SetStar(star)
	return &c
}
