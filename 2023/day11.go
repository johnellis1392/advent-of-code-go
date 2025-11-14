package aoc2023

import (
	"strings"

	common "github.com/johnellis1392/advent-of-code-go/common"
)

type Day11 struct {
	galaxies []common.Point
	nullRows []bool
	nullCols []bool
	width    int
	height   int
}

func (d *Day11) Galaxies() []common.Point {
	return d.galaxies
}

func (d *Day11) NullRows() []bool {
	return d.nullRows
}

func (d *Day11) NullCols() []bool {
	return d.nullCols
}

func (d *Day11) Width() int {
	return d.width
}

func (d *Day11) Height() int {
	return d.height
}

func (d *Day11) Year() string {
	return "2023"
}

func (d *Day11) Day() string {
	return "11"
}

func (d *Day11) Parse(input string) error {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	height, width := len(lines), len(lines[0])

	nullRows := make([]bool, height)
	for i := range nullRows {
		nullRows[i] = true
	}

	nullCols := make([]bool, width)
	for i := range nullCols {
		nullCols[i] = true
	}

	var galaxies []common.Point
	for row, line := range lines {
		line = strings.TrimSpace(line)
		for col, c := range line {
			if c == '#' {
				// Found a Galaxy
				nullRows[row] = false
				nullCols[col] = false
				galaxies = append(galaxies, common.PointFromRC(row, col))
			}
		}
	}

	d.galaxies = galaxies
	d.nullRows = nullRows
	d.nullCols = nullCols
	d.width = width
	d.height = height
	return nil
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// We just need to calculate the Manhattan distance and get the
// number of steps with the additional null-space increment.
func (d *Day11) PathLength(i, j int, nullSpaceSize int) uint64 {
	start, end := d.galaxies[i], d.galaxies[j]
	var path uint64 = 0
	for row := min(start.Y, end.Y) + 1; row <= max(start.Y, end.Y); row++ {
		if d.nullRows[row] {
			path += uint64(nullSpaceSize)
		} else {
			path += 1
		}
	}

	for col := min(start.X, end.X) + 1; col <= max(start.X, end.X); col++ {
		if d.nullCols[col] {
			path += uint64(nullSpaceSize)
		} else {
			path += 1
		}
	}

	return path
}

func (d *Day11) Part1() any {
	const nullSpaceSize = 2
	paths := make([][]uint64, len(d.galaxies))
	for i := 0; i < len(d.galaxies); i++ {
		paths[i] = make([]uint64, len(d.galaxies))
	}
	for i := 0; i < len(d.galaxies); i++ {
		paths[i][i] = 0
		for j := i + 1; j < len(d.galaxies); j++ {
			l := d.PathLength(i, j, nullSpaceSize)
			paths[i][j] = l
			paths[j][i] = l
		}
	}

	var res uint64 = 0
	for i := 0; i < len(d.galaxies)-1; i++ {
		for j := i + 1; j < len(d.galaxies); j++ {
			res += paths[i][j]
		}
	}

	return res
}

func (d *Day11) Part2() any {
	return 0
}
