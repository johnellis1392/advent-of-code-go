package aoc2023

import (
	"math"
	"strings"

	common "github.com/johnellis1392/advent-of-code-go/common"
)

type Day10 struct {
	input *Grid
}

func (d *Day10) Year() string {
	return "2023"
}

func (d *Day10) Day() string {
	return "10"
}

func (d *Day10) Parse(input string) error {
	var matrix [][]string
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		var row []string
		for i := 0; i < len(line); i++ {
			row = append(row, string(line[i]))
		}
		matrix = append(matrix, row)
	}
	d.input = NewGrid(matrix)

	return nil
}

type Grid struct {
	width, height int
	matrix        [][]string
	steps         [][]int
	start         common.Point
}

func (g *Grid) Size() (int, int) {
	return g.width, g.height
}

func (g *Grid) SetStep(p common.Point, i int) {
	g.steps[p.R()][p.C()] = i
}

func (g *Grid) GetStep(p common.Point) int {
	return g.steps[p.R()][p.C()]
}

func (g *Grid) Get(p common.Point) string {
	return g.matrix[p.R()][p.C()]
}

func (g *Grid) Reset() {
	width, height := g.Size()
	for r := 0; r < height; r++ {
		g.steps[r] = make([]int, width)
		for c := 0; c < width; c++ {
			g.steps[r][c] = math.MaxInt
		}
	}
	g.SetStep(g.start, 0)
}

func (g *Grid) Valid(p common.Point) bool {
	width, height := g.Size()
	return 0 <= p.R() && p.R() < height && 0 <= p.C() && p.C() < width
}

func (g *Grid) GetConnectedWalls(p common.Point) []common.Point {
	var res []common.Point
	r, c := p.R(), p.C()

	up := common.PointFromRC(r-1, c)
	down := common.PointFromRC(r+1, c)
	left := common.PointFromRC(r, c-1)
	right := common.PointFromRC(r, c+1)

	switch g.Get(p) {
	case "S":
		if g.Valid(up) {
			switch g.Get(up) {
			case "|", "7", "F":
				res = append(res, up)
			}
		}
		if g.Valid(down) {
			switch g.Get(down) {
			case "|", "J", "L":
				res = append(res, down)
			}
		}
		if g.Valid(left) {
			switch g.Get(left) {
			case "-", "L", "F":
				res = append(res, left)
			}
		}
		if g.Valid(right) {
			switch g.Get(right) {
			case "-", "J", "7":
				res = append(res, right)
			}
		}

	case "|":
		if g.Valid(up) {
			res = append(res, up)
		}
		if g.Valid(down) {
			res = append(res, down)
		}

	case "-":
		if g.Valid(left) {
			res = append(res, left)
		}
		if g.Valid(right) {
			res = append(res, right)
		}

	case "L":
		if g.Valid(up) {
			res = append(res, up)
		}
		if g.Valid(right) {
			res = append(res, right)
		}

	case "J":
		if g.Valid(up) {
			res = append(res, up)
		}
		if g.Valid(left) {
			res = append(res, left)
		}

	case "7":
		if g.Valid(left) {
			res = append(res, left)
		}
		if g.Valid(down) {
			res = append(res, down)
		}

	case "F":
		if g.Valid(right) {
			res = append(res, right)
		}
		if g.Valid(down) {
			res = append(res, down)
		}

	default:
		// No-op
	}
	return res
}

func (g *Grid) FewerSteps(from, to common.Point) bool {
	return g.GetStep(from)+1 < g.GetStep(to)
}

func (g *Grid) Step(from, to common.Point) {
	g.SetStep(to, g.GetStep(from)+1)
}

func (g *Grid) Adjacents(p common.Point) []common.Point {
	var res []common.Point
	up := common.PointFromRC(p.R()-1, p.C())
	if g.Valid(up) {
		res = append(res, up)
	}
	down := common.PointFromRC(p.R()+1, p.C())
	if g.Valid(down) {
		res = append(res, down)
	}
	left := common.PointFromRC(p.R(), p.C()-1)
	if g.Valid(left) {
		res = append(res, left)
	}
	right := common.PointFromRC(p.R(), p.C()+1)
	if g.Valid(right) {
		res = append(res, right)
	}
	return res
}

func NewGrid(matrix [][]string) *Grid {
	height, width := len(matrix), len(matrix[0])
	steps := make([][]int, height)
	var start common.Point

outer:
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if matrix[r][c] == "S" {
				start = common.PointFromRC(r, c)
				break outer
			}
		}
	}

	grid := &Grid{
		width:  width,
		height: height,
		matrix: matrix,
		steps:  steps,
		start:  start,
	}
	grid.Reset()
	return grid
}

type Queue struct {
	vs []common.Point
}

func (q *Queue) Enqueue(p common.Point) {
	q.vs = append(q.vs, p)
}

func (q *Queue) Pop() common.Point {
	v := q.vs[0]
	q.vs = q.vs[1:]
	return v
}

func (q *Queue) Empty() bool {
	return len(q.vs) == 0
}

func NewQueue() *Queue {
	return &Queue{[]common.Point{}}
}

type Set struct {
	vs map[common.Point]bool
}

func NewSet() *Set {
	return &Set{make(map[common.Point]bool)}
}

func (s *Set) First() common.Point {
	for p := range s.vs {
		return p
	}
	return common.PointFromRC(-1, -1)
}

func (s *Set) Add(p common.Point) {
	s.vs[p] = true
}

func (s *Set) Remove(p common.Point) {
	delete(s.vs, p)
}

func (s *Set) Contains(p common.Point) bool {
	_, ok := s.vs[p]
	return ok
}

func (s *Set) Diff(o *Set) (res *Set) {
	res = NewSet()
	s.ForEach(func(p common.Point) {
		if !o.Contains(p) {
			res.Add(p)
		}
	})
	return res
}

func (s *Set) Union(o *Set) (res *Set) {
	res = NewSet()
	s.ForEach(func(p common.Point) {
		res.Add(p)
	})
	o.ForEach(func(p common.Point) {
		res.Add(p)
	})
	return res
}

func (s *Set) Xor(o *Set) (res *Set) {
	res = NewSet()
	s.ForEach(func(p common.Point) {
		if !o.Contains(p) {
			res.Add(p)
		}
	})
	o.ForEach(func(p common.Point) {
		if !s.Contains(p) {
			res.Add(p)
		}
	})
	return res
}

func (s *Set) Intersect(o *Set) (res *Set) {
	res = NewSet()
	s.ForEach(func(p common.Point) {
		if o.Contains(p) {
			res.Add(p)
		}
	})
	return res
}

func (s *Set) Size() int {
	return len(s.vs)
}

func (s *Set) Overlaps(o *Set) bool {
	return s.Intersect(o).Size() > 0
}

func (s *Set) ForEach(f func(p common.Point)) {
	for p := range s.vs {
		f(p)
	}
}

func (d *Day10) Part1() any {
	grid := d.input
	frontier := NewQueue()
	frontier.Enqueue(grid.start)

	res := -1
	for !frontier.Empty() {
		currentPos := frontier.Pop()
		res = max(res, grid.GetStep(currentPos))
		for _, p := range grid.GetConnectedWalls(currentPos) {
			if grid.FewerSteps(currentPos, p) {
				frontier.Enqueue(p)
				grid.Step(currentPos, p)
			}
		}
	}

	return res
}

func getWalls(grid *Grid) *Set {
	frontier := NewQueue()
	frontier.Enqueue(grid.start)
	walls := NewSet()

	for !frontier.Empty() {
		currentPos := frontier.Pop()
		for _, p := range grid.GetConnectedWalls(currentPos) {
			if !walls.Contains(p) {
				walls.Add(p)
				frontier.Enqueue(p)
			}
		}
	}

	return walls
}

func getTiles(grid *Grid, walls *Set) *Set {
	w, h := grid.Size()
	tiles := NewSet()
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			p := common.PointFromRC(r, c)
			if !walls.Contains(p) {
				tiles.Add(p)
			}
		}
	}
	return tiles
}

func contained(grid *Grid, walls *Set, p common.Point) bool {
	res := 0

	_, h := grid.Size()
	r := p.R()
	for ; r < h; r++ {
		if !walls.Contains(common.PointFromRC(r, p.C())) {
			continue
		}
		s := grid.Get(common.PointFromRC(r, p.C()))
		switch s {
		case "-":
			res++
		case "F":
			res++
			r++
			for ; r < h; r++ {
				s2 := grid.Get(common.PointFromRC(r, p.C()))
				if s2 == "J" {
					break
				} else if s2 == "L" {
					res++
					break
				}
			}
		case "7":
			res++
			r++
			for ; r < h; r++ {
				s2 := grid.Get(common.PointFromRC(r, p.C()))
				if s2 == "L" {
					break
				} else if s2 == "J" {
					res++
					break
				}
			}
		}
	}

	return res%2 == 1
}

func replaceStart(grid *Grid) {
	p := grid.start
	ps := grid.GetConnectedWalls(p)
	a, b := ps[0], ps[1]

	eq := func(q1, q2 common.Point) bool {
		return a == q1 && b == q2 || b == q2 && a == q1
	}

	var s string
	switch {
	case eq(p.South(), p.North()):
		s = "|"
	case eq(p.West(), p.East()):
		s = "-"
	case eq(p.West(), p.North()):
		s = "J"
	case eq(p.North(), p.East()):
		s = "L"
	case eq(p.East(), p.South()):
		s = "F"
	case eq(p.South(), p.West()):
		s = "7"
	}
	grid.matrix[p.R()][p.C()] = s
}

func (d *Day10) Part2() any {
	grid := d.input
	walls := getWalls(grid)
	tiles := getTiles(grid, walls)
	replaceStart(grid)

	sum := 0
	containedPoints := NewSet()
	tiles.ForEach(func(p common.Point) {
		if contained(grid, walls, p) {
			sum++
			containedPoints.Add(p)
		}
	})

	return sum
}
