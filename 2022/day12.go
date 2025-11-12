package aoc2022

import (
	"math"
	"strings"

	common "github.com/johnellis1392/advent-of-code-go/common"
)

type Day12 struct {
	input string
}

func (d *Day12) Year() string {
	return "2022"
}

func (d *Day12) Day() string {
	return "12"
}

func (d *Day12) Parse(input string) error {
	d.input = input
	return nil
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}

func height(c byte) int {
	return int(c - 'a')
}

type Grid struct {
	matrix           [][]string
	heightMap        [][]int
	steps            [][]int
	start, end       common.Point
	numRows, numCols int
}

func NewGrid(input [][]byte) *Grid {
	m, n := len(input), len(input[0])
	matrix := make([][]string, m)
	heightMap := make([][]int, m)
	steps := make([][]int, m)
	var start, end common.Point

	for r := 0; r < m; r++ {
		matrix[r] = make([]string, n)
		heightMap[r] = make([]int, n)
		steps[r] = make([]int, n)
		for c := 0; c < n; c++ {
			matrix[r][c] = string(input[r][c])
			steps[r][c] = math.MaxInt
			switch input[r][c] {
			case 'S':
				heightMap[r][c] = height('a')
				start = common.PointFromRC(r, c)
			case 'E':
				heightMap[r][c] = height('z')
				end = common.PointFromRC(r, c)
			default:
				heightMap[r][c] = height(input[r][c])
			}
		}
	}

	return &Grid{
		matrix:    matrix,
		heightMap: heightMap,
		steps:     steps,
		start:     start,
		end:       end,
		numRows:   m,
		numCols:   n,
	}
}

func (g *Grid) Get(p common.Point) (string, int) {
	return g.matrix[p.R()][p.C()], g.heightMap[p.R()][p.C()]
}

func (g *Grid) Size() (int, int) {
	return g.numRows, g.numCols
}

func (g *Grid) Reset() {
	m, n := g.Size()
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			g.steps[r][c] = math.MaxInt
		}
	}
}

func (g *Grid) StartingPoints() []common.Point {
	var res []common.Point
	m, n := g.Size()
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if g.matrix[r][c] == "a" {
				res = append(res, common.PointFromRC(r, c))
			}
		}
	}
	return res
}

func (g *Grid) Adjacents(p common.Point) []common.Point {
	var adjacents []common.Point
	numRows, numCols := g.Size()
	if p.R()-1 >= 0 {
		adjacents = append(adjacents, common.PointFromRC(p.R()-1, p.C()))
	}
	if p.R()+1 < numRows {
		adjacents = append(adjacents, common.PointFromRC(p.R()+1, p.C()))
	}
	if p.C()-1 >= 0 {
		adjacents = append(adjacents, common.PointFromRC(p.R(), p.C()-1))
	}
	if p.C()+1 < numCols {
		adjacents = append(adjacents, common.PointFromRC(p.R(), p.C()+1))
	}
	return adjacents
}

func (g *Grid) Moveable(from, to common.Point) bool {
	_, h1 := g.Get(from)
	_, h2 := g.Get(to)
	if h2 <= h1 || h2-h1 == 1 {
		return true
	} else {
		return false
	}
}

func (g *Grid) FewerSteps(from, to common.Point) bool {
	s1 := g.GetStep(from)
	s2 := g.GetStep(to)
	if s1+1 < s2 {
		return true
	} else {
		return false
	}
}

func (g *Grid) GetStep(p common.Point) int {
	return g.steps[p.R()][p.C()]
}

func (g *Grid) SetStep(p common.Point, i int) {
	g.steps[p.R()][p.C()] = i
}

func (g *Grid) Step(from, to common.Point) {
	g.SetStep(to, g.GetStep(from)+1)
}

func readInput(input string) *Grid {
	lines := strings.Split(input, "\n")
	var res [][]byte
	for r := 0; r < len(lines); r++ {
		line := strings.TrimSpace(lines[r])
		if len(line) == 0 {
			continue
		}
		res = append(res, make([]byte, len(line)))
		for c := 0; c < len(line); c++ {
			res[r][c] = line[c]
		}
	}
	return NewGrid(res)
}

type Queue struct {
	q []common.Point
}

func NewQueue() *Queue {
	return &Queue{[]common.Point{}}
}

func (q *Queue) Enqueue(p common.Point) {
	q.q = append(q.q, p)
}

func (q *Queue) Pop() common.Point {
	p := q.q[0]
	q.q = q.q[1:]
	return p
}

func (q *Queue) Empty() bool {
	return len(q.q) == 0
}

func shortestPath(grid *Grid, start, end common.Point) int {
	grid.SetStep(start, 0)

	frontier := NewQueue()
	frontier.Enqueue(start)

	for !frontier.Empty() {
		currentPos := frontier.Pop()
		for _, p := range grid.Adjacents(currentPos) {
			if grid.Moveable(currentPos, p) && grid.FewerSteps(currentPos, p) {
				grid.Step(currentPos, p)
				frontier.Enqueue(p)
			}
		}
	}

	return grid.GetStep(end)
}

func (d *Day12) Part1() any {
	grid := readInput(d.input)
	start, end := grid.start, grid.end
	grid.Reset()
	return shortestPath(grid, start, end)
}

func (d *Day12) Part2() any {
	grid := readInput(d.input)
	end := grid.end
	n := math.MaxInt
	for _, start := range grid.StartingPoints() {
		n = min(n, shortestPath(grid, start, end))
	}
	return n
}
