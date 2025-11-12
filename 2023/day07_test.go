package aoc2023_test

import (
	"testing"

	aoc "github.com/johnellis1392/advent-of-code-go/2023"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay07_Part1(t *testing.T) {
	input := `32T3K 765
	T55J5 684
	KK677 28
	KTJJT 220
	QQQJA 483`

	expected := 0
	day := new(aoc.Day07)
	util.TestPart1(t, day, input, expected)
}
