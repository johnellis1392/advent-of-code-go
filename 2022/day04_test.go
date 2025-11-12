package aoc2022_test

import (
	"testing"

	aoc2022 "github.com/johnellis1392/advent-of-code-go/2022"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay04_Part1(t *testing.T) {
	input :=
		`2-4,6-8
	   2-3,4-5
	   5-7,7-9
	   2-8,3-7
	   6-6,4-6
	   2-6,4-8`
	expected := 0
	day := new(aoc2022.Day04)
	util.TestPart1(t, day, input, expected)
}
