package aoc2023_test

import (
	aoc "github.com/johnellis1392/advent-of-code-go/2023"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
	"testing"
)

func TestDay06_Part1(t *testing.T) {
	input := `Time:      7  15   30
	Distance:  9  40  200`
	expected := 0
	day := new(aoc.Day06)
	util.TestPart1(t, day, input, expected)
}
