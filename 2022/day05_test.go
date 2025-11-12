package aoc2022_test

import (
	"testing"

	aoc2022 "github.com/johnellis1392/advent-of-code-go/2022"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay05_Part1(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	expected := 0
	day := new(aoc2022.Day05)
	util.TestPart1(t, day, input, expected)
}
