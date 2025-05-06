package vector

import (
	"github.com/GeekchanskiY/avagen/pkg/point"
)

type Vector struct {
	X, Y int
}

func (v *Vector) FromPoints(p1, p2 point.Point) {
	v.X = p1.X - p2.X
	v.Y = p1.Y - p2.Y
}

func Edge(v0, v1, p point.Point) int {
	var (
		a, b Vector
	)

	a.FromPoints(p, v1)
	b.FromPoints(v1, v0)

	return a.X*b.Y - a.Y*b.X
}
