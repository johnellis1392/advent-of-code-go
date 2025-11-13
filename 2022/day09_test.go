package aoc2022_test

import (
	"testing"

	aoc2022 "github.com/johnellis1392/advent-of-code-go/2022"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay09_Part1(t *testing.T) {
	input := `R 4
	U 4
	L 3
	D 1
	R 4
	D 1
	L 5
	R 2`
	// input := `R 5
	// U 8
	// L 8
	// D 3
	// R 17
	// D 10
	// L 25
	// U 20`

	expected := 13
	day := new(aoc2022.Day09)
	util.TestPart1(t, day, input, expected)
}
