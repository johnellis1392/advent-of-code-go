package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	r, c int
}

func (p Point) String() string {
	return fmt.Sprintf("(r=%d, c=%d)", p.r, p.c)
}

func (p Point) South() Point {
	return Point{p.r + 1, p.c}
}

func (p Point) North() Point {
	return Point{p.r - 1, p.c}
}

func (p Point) West() Point {
	return Point{p.r, p.c - 1}
}

func (p Point) East() Point {
	return Point{p.r, p.c + 1}
}

type Grid struct {
	width, height int
	start, end    Point
	matrix        [][]string
}

func NewGrid(p0, p1 Point, padFloor bool) *Grid {
	width := p1.c - p0.c + 1
	height := p1.r - p0.r + 1
	if padFloor {
		height += 2
	}
	matrix := make([][]string, height)
	for r := 0; r < height; r++ {
		matrix[r] = make([]string, width)
		for c := 0; c < width; c++ {
			matrix[r][c] = "."
		}
	}
	return &Grid{
		width:  width,
		height: height,
		start:  p0,
		end:    p1,
		matrix: matrix,
	}
}

func (g *Grid) Set(p Point, v string) {
	g.matrix[p.r-g.start.r][p.c-g.start.c] = v
}

func (g *Grid) Get(p Point) string {
	return g.matrix[p.r-g.start.r][p.c-g.start.c]
}

func (g *Grid) Valid(p Point) bool {
	return p.r < g.height && g.start.c <= p.c && p.c <= g.end.c
}

func (g *Grid) IsBlock(p Point) bool {
	switch g.Get(p) {
	case "#", "o":
		return true
	default:
		return false
	}
}

func (g *Grid) DrawLine(from, to Point) {
	r1, r2 := min(from.r, to.r), max(from.r, to.r)
	c1, c2 := min(from.c, to.c), max(from.c, to.c)
	for r := r1; r <= r2; r++ {
		for c := c1; c <= c2; c++ {
			g.Set(Point{r, c}, "#")
		}
	}
}

func (g *Grid) Size() (int, int) {
	return g.width, g.height
}

func (g *Grid) Dump() {
	w, h := g.Size()
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			fmt.Print(g.matrix[r][c])
		}
		fmt.Println()
	}
}

func readInput(input string, padFloor bool) *Grid {
	var c1, c2, height int
	c1 = 500
	c2 = 500
	height = 0

	var pointConfigs [][]Point
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		var points []Point
		for _, ps := range strings.Split(line, " -> ") {
			coords := strings.Split(ps, ",")
			c, _ := strconv.Atoi(coords[0])
			r, _ := strconv.Atoi(coords[1])
			p := Point{r, c}
			points = append(points, p)
			if c < c1 {
				c1 = c
			}
			if c > c2 {
				c2 = c
			}
			if r > height {
				height = r
			}
		}
		pointConfigs = append(pointConfigs, points)
	}

	start := Point{0, c1}
	end := Point{height, c2}
	grid := NewGrid(start, end, padFloor)

	for _, points := range pointConfigs {
		for i := 1; i < len(points); i++ {
			grid.DrawLine(points[i-1], points[i])
		}
	}

	if padFloor {
		grid.DrawLine(Point{height + 2, start.c}, Point{height + 2, end.c})
	}

	grid.Set(Point{0, 500}, "+")
	return grid
}

func part1(input string) int {
	grid := readInput(input, false)
	dropPoint := Point{0, 500}

	res := 0
	p := dropPoint
	for true {
		if !grid.IsBlock(p.South()) {
			p = p.South()
		} else if !grid.Valid(p.South().West()) {
			break
		} else if !grid.IsBlock(p.South().West()) {
			p = p.South().West()
		} else if !grid.Valid(p.South().East()) {
			break
		} else if !grid.IsBlock(p.South().East()) {
			p = p.South().East()
		} else {
			grid.Set(p, "o")
			p = dropPoint
			res++
		}
	}

	return res
}

func part2(input string) int {
	grid := readInput(input, true)
	grid.Dump()
	return 0
}

func main() {
	const DEBUG = true
	filename := "input.txt"
	test_input := `498,4 -> 498,6 -> 496,6
	503,4 -> 502,4 -> 502,9 -> 494,9
	`

	var input string
	if DEBUG {
		input = test_input
	} else {
		s, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		input = string(s)
	}

	fmt.Printf("2022 Day 14, Part 1: %v\n", part1(input))
	fmt.Printf("2022 Day 14, Part 2: %v\n", part2(input))
}
