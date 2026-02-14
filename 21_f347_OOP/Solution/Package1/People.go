package package1

import (
	"fmt"
	"time"
)

type People struct {
	FirstName string
	LastName  string
	birthday  time.Time
}

func (p *People) SetBirthday(birthday time.Time) {
	if birthday.Before(time.Now()) {
		p.birthday = birthday
	} else {
		panic("Birthday is not more than a present day")
	}
}

func (p *People) GetBirthday() time.Time {
	return p.birthday
}

func (p *People) getfullname() string {
	return p.FirstName + " " + p.LastName
}

func (p *People) GetAllinString() string {
	return fmt.Sprintf("People\nFullName: %v\nBirthday: %v\n\n", p.getfullname(), p.GetBirthday().String())
}

func NewPeople(firstname string, lastname string, birthday time.Time) *People {
	p := People{FirstName: firstname, LastName: lastname}
	p.SetBirthday(birthday)
	return &p
}
