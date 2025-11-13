package aoc2022_test

import (
	"testing"

	aoc2022 "github.com/johnellis1392/advent-of-code-go/2022"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay08_Part1(t *testing.T) {
	input := `30373
	25512
	65332
	33549
	35390`

	expected := 21
	day := new(aoc2022.Day08)
	util.TestPart1(t, day, input, expected)
}
