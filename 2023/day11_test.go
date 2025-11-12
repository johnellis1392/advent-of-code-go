package aoc2023_test

import (
	"testing"

	aoc "github.com/johnellis1392/advent-of-code-go/2023"
	tu "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay11_Part1(t *testing.T) {
	d := new(aoc.Day11)
	input := `...#......
	.......#..
	#.........
	..........
	......#...
	.#........
	.........#
	..........
	.......#..
	#...#.....`
	expected := 0
	tu.TestPart1(t, d, input, expected)
}
