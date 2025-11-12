package aoc2022

import (
	"strconv"
	"strings"

	common "github.com/johnellis1392/advent-of-code-go/common"
)

type Day09 struct {
	input string
}

func (d *Day09) Year() string {
	return "2022"
}

func (d *Day09) Day() string {
	return "09"
}

func (d *Day09) Parse(input string) error {
	d.input = input
	return nil
}

func normalize(x int) int {
	if x == 0 {
		return 0
	} else if x < 0 {
		return -1
	} else {
		return 1
	}
}

func (d *Day09) Part1() any {
	head := common.Point{X: 0, Y: 0}
	tail := common.Point{X: 0, Y: 0}
	var lastPos common.Point

	steps := make(map[common.Point]bool)
	steps[tail] = true

	for _, line := range strings.Split(d.input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		move := strings.Split(line, " ")
		dir := move[0]
		n, _ := strconv.Atoi(move[1])

		var dx, dy int
		switch dir {
		case "R":
			dx, dy = 1, 0
		case "U":
			dx, dy = 0, 1
		case "L":
			dx, dy = -1, 0
		case "D":
			dx, dy = 0, -1
		}

		for i := 0; i < n; i++ {
			lastPos = head
			head.X += dx
			head.Y += dy
			if abs(head.X-tail.X) > 1 || abs(head.Y-tail.Y) > 1 {
				tail = lastPos
				if _, ok := steps[tail]; !ok {
					steps[tail] = true
				}
			}
		}
	}

	return len(steps)
}

// func dump(head Point, knots []Point, steps map[Point]bool, x1, y1, x2, y2 int) {
// 	w, h := x2-x1+2, y2-y1+2
// 	graph := make([][]string, h)
// 	for i := 0; i < h; i++ {
// 		graph[i] = make([]string, w)
// 	}

// 	for r := 0; r < h; r++ {
// 		for c := 0; c < w; c++ {
// 			graph[r][c] = "."
// 		}
// 	}

// 	for step, _ := range steps {
// 		x, y := step.x-x1+1, step.y-y1+1
// 		graph[y][x] = "#"
// 	}

// 	graph[-y1+1][-x1+1] = "s"

// 	for i := len(knots) - 1; i >= 0; i-- {
// 		point := knots[i]
// 		x, y := point.x-x1+1, point.y-y1+1
// 		graph[y][x] = fmt.Sprintf("%d", i+1)
// 	}

// 	graph[head.y-y1+1][head.x-x1+1] = "H"

// 	for r := h - 1; r >= 0; r-- {
// 		for c := 0; c < w; c++ {
// 			fmt.Printf("%s", graph[r][c])
// 		}
// 		fmt.Println()
// 	}
// }

func (d *Day09) Part2() any {
	x1, y1, x2, y2 := 0, 0, 0, 0
	head := common.Point{X: 0, Y: 0}
	knots := make([]common.Point, 9)
	for i := 0; i < 9; i++ {
		knots[i] = common.Point{X: 0, Y: 0}
	}
	steps := make(map[common.Point]bool)
	steps[head] = true

	for _, line := range strings.Split(d.input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		move := strings.Split(line, " ")
		dir := move[0]
		n, _ := strconv.Atoi(move[1])

		var dx, dy int
		switch dir {
		case "R":
			dx, dy = 1, 0
		case "L":
			dx, dy = -1, 0
		case "U":
			dx, dy = 0, 1
		case "D":
			dx, dy = 0, -1
		}

		for i := 0; i < n; i++ {
			head.X += dx
			head.Y += dy
			curr := head

			for j := 0; j < 9; j++ {
				if abs(curr.X-knots[j].X) <= 1 && abs(curr.Y-knots[j].Y) <= 1 {
					break
				}

				ddx, ddy := curr.X-knots[j].X, curr.Y-knots[j].Y
				ddx, ddy = normalize(ddx), normalize(ddy)
				knots[j].X += ddx
				knots[j].Y += ddy

				curr = knots[j]
			}

			if _, ok := steps[knots[9-1]]; !ok {
				steps[knots[9-1]] = true
			}

			x1 = min(x1, head.X)
			x2 = max(x2, head.X)
			y1 = min(y1, head.Y)
			y2 = max(y2, head.Y)
		}
	}

	return len(steps)
}
