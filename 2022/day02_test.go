package aoc2022_test

import (
	"testing"

	aoc2022 "github.com/johnellis1392/advent-of-code-go/2022"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay02_Part1(t *testing.T) {
	day := new(aoc2022.Day02)
	input := `A Y
B X
C Z`
	expected := 15
	util.TestPart1(t, day, input, expected)
}
