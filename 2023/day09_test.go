package aoc2023_test

import (
	aoc "github.com/johnellis1392/advent-of-code-go/2023"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
	"testing"
)

func TestDay09_Part1(t *testing.T) {
	input := `0 3 6 9 12 15
	1 3 6 10 15 21
	10 13 16 21 30 45`
	expected := 0
	day := new(aoc.Day09)
	util.TestPart1(t, day, input, expected)
}
