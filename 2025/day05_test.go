package aoc2025_test

import (
	"testing"

	aoc "github.com/johnellis1392/advent-of-code-go/2025"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

const input = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestDay05_Part1(t *testing.T) {
	expected := 3
	day := new(aoc.Day05)
	util.TestPart1(t, day, input, expected)
}

func TestDay05_Part2(t *testing.T) {
	expected := int64(14)
	day := new(aoc.Day05)
	util.TestPart2(t, day, input, expected)
}
