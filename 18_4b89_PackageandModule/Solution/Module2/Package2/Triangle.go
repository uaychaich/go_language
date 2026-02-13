package package2

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.2 * t.Base * t.Height
}
