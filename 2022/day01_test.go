package aoc2022_test

import (
	"testing"

	aoc2022 "github.com/johnellis1392/advent-of-code-go/2022"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay01_Part1(t *testing.T) {
	day := new(aoc2022.Day01)
	input := `1000
	2000
	3000
	
	4000
	
	5000
	6000
	
	7000
	8000
	9000
	
	10000`
	expected := 24000
	util.TestPart1(t, day, input, expected)
}
