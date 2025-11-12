package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	r, c int
}

type Grid struct {
	width, height int
	matrix        [][]string
	galaxies      []Point
}

func NewGrid(matrix [][]string) *Grid {
	height := len(matrix)
	width := len(matrix[0])

	var galaxies []Point
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if matrix[r][c] == "#" {
				galaxies = append(galaxies, Point{r, c})
			}
		}
	}

	return &Grid{
		matrix:   matrix,
		width:    width,
		height:   height,
		galaxies: galaxies,
	}
}

func makeNullRow(width int) []string {
	row := make([]string, width)
	for i := 0; i < width; i++ {
		row[i] = "."
	}
	return row
}

func readInput(input string) *Grid {
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
loop1:
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if ss[r][c] == "#" {
				continue loop1
			}
		}
		nullRows[r] = true
	}

	// Check Cols
	nullCols := make(map[int]bool)
loop2:
	for c := 0; c < w; c++ {
		for r := 0; r < h; r++ {
			if ss[r][c] == "#" {
				continue loop2
			}
		}
		nullCols[c] = true
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

	grid := NewGrid(matrix)
	return grid
}

func (grid *Grid) Size() (int, int) {
	return grid.width, grid.height
}

func (grid *Grid) Dump() {
	w, h := grid.Size()
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			fmt.Print(grid.matrix[r][c])
		}
		fmt.Println()
	}
}

func mag(i int) int {
	switch {
	case i > 0:
		return 1
	case i < 0:
		return -1
	default:
		return 0
	}
}

func abs(i int) int {
	switch {
	case i > 0:
		return i
	case i < 0:
		return i * -1
	default:
		return 0
	}
}

func shortestPathSize(from, to Point) int {
	res := 0
	current := from
	for current != to {
		res++
		dr := to.r - current.r
		dc := to.c - current.c
		if abs(dr) > abs(dc) {
			current = Point{current.r + mag(dr), current.c}
		} else {
			current = Point{current.r, current.c + mag(dc)}
		}
	}
	return res
}

func part1(input string) int {
	grid := readInput(input)
	// grid.Dump()
	sum := 0
	for i := 0; i < len(grid.galaxies)-1; i++ {
		for j := i + 1; j < len(grid.galaxies); j++ {
			a, b := grid.galaxies[i], grid.galaxies[j]
			sum += shortestPathSize(a, b)
		}
	}
	return sum
}

func part2(input string) int {
	return 0
}

func main() {
	const DEBUG = false
	filename := "input.txt"
	test_input := `...#......
	.......#..
	#.........
	..........
	......#...
	.#........
	.........#
	..........
	.......#..
	#...#.....`

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

	fmt.Printf("2023 Day 11, Part 1: %v\n", part1(input))
	fmt.Printf("2023 Day 11, Part 2: %v\n", part2(input))
}
