package aoc2022

import (
	"strconv"
	"strings"

	common "github.com/johnellis1392/advent-of-code-go/common"
)

type Day14 struct {
	grid *common.Grid
}

func (d *Day14) Year() string {
	return "2022"
}

func (d *Day14) Day() string {
	return "14"
}

func (d *Day14) Parse(input string) error {
	var c1, c2, height int
	c1 = 500
	c2 = 500
	height = 0

	var pointConfigs [][]common.Point
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		var points []common.Point
		for _, ps := range strings.Split(line, " -> ") {
			coords := strings.Split(ps, ",")
			c, _ := strconv.Atoi(coords[0])
			r, _ := strconv.Atoi(coords[1])
			p := common.PointFromRC(r, c)
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

	start := common.PointFromRC(0, c1)
	end := common.PointFromRC(height, c2)
	grid := common.NewGridFromPoints(start, end)
	d.grid = grid

	for _, points := range pointConfigs {
		for i := 1; i < len(points); i++ {
			d.DrawLine(points[i-1], points[i])
		}
	}

	d.grid.Set(common.PointFromRC(0, 500), "+")
	return nil
}

func (d *Day14) padFloor() {
	from := common.PointFromRC(d.grid.Height+2, d.grid.Start.C())
	to := common.PointFromRC(d.grid.Height+2, d.grid.End.C())
	d.DrawLine(from, to)
}

func (d *Day14) IsBlock(p common.Point) bool {
	switch *d.grid.Get(p) {
	case "#", "o":
		return true
	default:
		return false
	}
}

func (d *Day14) DrawLine(from, to common.Point) {
	r1, r2 := min(from.R(), to.R()), max(from.R(), to.R())
	c1, c2 := min(from.C(), to.C()), max(from.C(), to.C())
	for r := r1; r <= r2; r++ {
		for c := c1; c <= c2; c++ {
			d.grid.Set(common.PointFromRC(r, c), "#")
		}
	}
}

func (d *Day14) Part1() any {
	dropPoint := common.PointFromRC(0, 500)

	res := 0
	p := dropPoint
	for {
		if !d.IsBlock(p.South()) {
			p = p.South()
		} else if !d.grid.Contains(p.South().West()) {
			break
		} else if !d.IsBlock(p.South().West()) {
			p = p.South().West()
		} else if !d.grid.Contains(p.South().East()) {
			break
		} else if !d.IsBlock(p.South().East()) {
			p = p.South().East()
		} else {
			d.grid.Set(p, "o")
			p = dropPoint
			res++
		}
	}

	return res
}

func (d *Day14) Part2() any {
	return 0
}
