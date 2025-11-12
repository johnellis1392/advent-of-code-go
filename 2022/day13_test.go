package aoc2022_test

import (
	"testing"

	aoc2022 "github.com/johnellis1392/advent-of-code-go/2022"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay13_Part1(t *testing.T) {
	input := `[1,1,3,1,1]
	[1,1,5,1,1]
	
	[[1],[2,3,4]]
	[[1],4]
	
	[9]
	[[8,7,6]]
	
	[[4,4],4,4]
	[[4,4],4,4,4]
	
	[7,7,7,7]
	[7,7,7]
	
	[]
	[3]
	
	[[[]]]
	[[]]
	
	[1,[2,[3,[4,[5,6,7]]]],8,9]
	[1,[2,[3,[4,[5,6,0]]]],8,9]`

	expected := 0
	day := new(aoc2022.Day13)
	util.TestPart1(t, day, input, expected)
}
