package package2

type People struct {
	Firstname string
	Lastname  string
}

func (p People) getFullname() string {
	return p.Firstname + " " + p.Lastname
}
