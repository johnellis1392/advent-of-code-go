package common

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func PointFromRC(r, c int) Point {
	return Point{X: c, Y: r}
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p Point) R() int {
	return p.Y
}

func (p Point) C() int {
	return p.X
}

func (p Point) North() Point {
	return PointFromRC(p.R()-1, p.C())
}

func (p Point) South() Point {
	return PointFromRC(p.R()+1, p.C())
}

func (p Point) East() Point {
	return PointFromRC(p.R(), p.C()+1)
}

func (p Point) West() Point {
	return PointFromRC(p.R(), p.C()-1)
}

func (p Point) String() string {
	return fmt.Sprintf("(r=%d, c=%d)", p.R(), p.C())
}

func (p Point) Within(start, end Point) bool {
	return start.X <= p.X && p.X <= end.X && start.Y <= p.Y && p.Y <= end.Y
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	} else {
		return x
	}
}

func Mag(i int) int {
	switch {
	case i > 0:
		return 1
	case i < 0:
		return -1
	default:
		return 0
	}
}
