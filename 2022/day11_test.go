package aoc2022_test

import (
	"testing"

	aoc2022 "github.com/johnellis1392/advent-of-code-go/2022"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay11_Part1(t *testing.T) {
	input := ""
	expected := 0
	day := new(aoc2022.Day11)
	util.TestPart1(t, day, input, expected)
}
