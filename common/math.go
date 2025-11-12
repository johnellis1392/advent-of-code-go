package common

type Point struct {
	X, Y int
}

func PointFromRC(r, c int) Point {
	return Point{X: c, Y: r}
}

func (p Point) R() int {
	return p.Y
}

func (p Point) C() int {
	return p.X
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	} else {
		return x
	}
}
