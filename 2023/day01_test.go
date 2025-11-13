package aoc2023_test

import (
	"testing"

	aoc "github.com/johnellis1392/advent-of-code-go/2023"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

// input := `two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen`
func TestDay01_Part1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	expected := 142
	day := new(aoc.Day01)
	util.TestPart1(t, day, input, expected)
}
