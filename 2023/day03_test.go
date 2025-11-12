package aoc2023_test

import (
	"testing"

	aoc "github.com/johnellis1392/advent-of-code-go/2023"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay03_Part1(t *testing.T) {
	input :=
		`467..114..
		 ...*......
		 ..35..633.
		 ......#...
		 617*......
		 .....+.58.
		 ..592.....
		 ......755.
		 ...$.*....
		 .664.598..`

	expected := 0
	day := new(aoc.Day03)
	util.TestPart1(t, day, input, expected)
}
