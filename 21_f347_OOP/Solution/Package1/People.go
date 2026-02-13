package package1

import "time"

type People struct {
	FirstName string
	LastName  string
	birthday  time.Time
}

func (p People) SetBirthday(birthday time.Time) {
	if birthday.Before(time.Now()) {
		p.birthday = birthday
	} else {
		panic("Birthday is not more than a present day")
	}
}

func (p People) GetBirthday() time.Time {
	return p.birthday
}

func (p People) getfullname() string {
	return p.FirstName + " " + p.LastName
}

func (p People) GetAllinString() string {
	return p.getfullname() + " " + p.GetBirthday().Format("2006-01-02 15:04:05")
}

func (p People) New(firstname string, lastname string, birthday time.Time) {
	p.FirstName = firstname
	p.LastName = lastname
	p.SetBirthday(birthday)
}
