package aoc2022_test

import (
	"testing"

	aoc "github.com/johnellis1392/advent-of-code-go/2022"
	tu "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay14_Part1(t *testing.T) {
	d := new(aoc.Day14)
	input := `498,4 -> 498,6 -> 496,6
	503,4 -> 502,4 -> 502,9 -> 494,9
	`
	expected := 0
	tu.TestPart1(t, d, input, expected)
}
