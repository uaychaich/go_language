package package1

import "fmt"

type Animal struct {
	Height float64
	Weight float64
}

func (a Animal) getHeightWeight() string {
	return fmt.Sprintf("Height=%v | Weight=%v", a.Height, a.Weight)
}
