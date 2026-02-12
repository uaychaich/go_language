package main

type twonums struct {
	x int
	y int
}

func (tn twonums) plus() int {
	return tn.x + tn.y
}

func (tn twonums) sub() int {
	return tn.x - tn.y
}
