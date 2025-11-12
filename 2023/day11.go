package aoc2023

import (
	"strings"

	common "github.com/johnellis1392/advent-of-code-go/common"
)

type Day11 struct {
	grid     *common.Grid
	galaxies []common.Point
}

func (d *Day11) Year() string {
	return "2023"
}

func (d *Day11) Day() string {
	return "11"
}

func makeNullRow(width int) []string {
	row := make([]string, width)
	for i := 0; i < width; i++ {
		row[i] = "."
	}
	return row
}

func (d *Day11) Parse(input string) error {
	var ss [][]string
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		var row []string
		for i := 0; i < len(line); i++ {
			row = append(row, string(line[i]))
		}
		ss = append(ss, row)
	}

	w, h := len(ss[0]), len(ss)

	// Check Rows
	nullRows := make(map[int]bool)
	nullCols := make(map[int]bool)
	var galaxies []common.Point
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if ss[r][c] == "#" {
				nullRows[r] = true
				nullCols[c] = true
				galaxies = append(galaxies, common.PointFromRC(r, c))
			}
		}
	}

	newWidth := w + len(nullCols)*2
	var matrix [][]string
	for r := 0; r < h; r++ {
		if _, isNullRow := nullRows[r]; isNullRow {
			matrix = append(matrix, makeNullRow(newWidth))
			matrix = append(matrix, makeNullRow(newWidth))
			continue
		}
		var row []string
		for c := 0; c < w; c++ {
			if _, isNullCol := nullCols[c]; isNullCol {
				row = append(row, ".", ".")
				continue
			}
			row = append(row, ss[r][c])
		}
		matrix = append(matrix, row)
	}

	d.galaxies = galaxies
	d.grid = nil // TODO: Fix this

	return nil
}

func shortestPathSize(from, to common.Point) int {
	res := 0
	current := from
	for current != to {
		res++
		dr := to.R() - current.R()
		dc := to.C() - current.C()
		if common.Abs(dr) > common.Abs(dc) {
			current = common.PointFromRC(current.R()+common.Mag(dr), current.C())
		} else {
			current = common.PointFromRC(current.R(), current.C()+common.Mag(dc))
		}
	}
	return res
}

func (d *Day11) Part1() any {
	sum := 0
	for i := 0; i < len(d.galaxies)-1; i++ {
		for j := i + 1; j < len(d.galaxies); j++ {
			a, b := d.galaxies[i], d.galaxies[j]
			sum += shortestPathSize(a, b)
		}
	}
	return sum
}

func (d *Day11) Part2() any {
	return 0
}
