package shape

import "math"

type Circle struct {
	R float32
}

func (c *Circle) Area() float32 {
	return 2 * math.Pi * c.R * c.R
}
func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.R
}
