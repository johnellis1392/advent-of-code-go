package aoc2022_test

import (
	"testing"

	aoc2022 "github.com/johnellis1392/advent-of-code-go/2022"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay12_Part1(t *testing.T) {
	input := `Sabqponm
	abcryxxl
	accszExk
	acctuvwj
	abdefghi`
	expected := 31
	day := new(aoc2022.Day12)
	util.TestPart1(t, day, input, expected)
}
