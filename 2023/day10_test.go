package aoc2023_test

import (
	"testing"

	aoc "github.com/johnellis1392/advent-of-code-go/2023"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

// test_input =
// 	`FF7FSF7F7F7F7F7F---7
// 	 L|LJ||||||||||||F--J
// 	 FL-7LJLJ||||||LJL-77
// 	 F--JF--7||LJLJ7F7FJ-
// 	 L---JF-JLJ.||-FJLJJ7
// 	 |F|F-JF---7F7-L7L|7|
// 	 |FFJF7L7F-JF7|JL---7
// 	 7-L-JL7||F7|L7F-7F7|
// 	 L.L7LFJ|||||FJL7||LJ
// 	 L7JLJL-JLJLJL--JLJ.L`

func TestDay10_Part1(t *testing.T) {
	input :=
		`.....
	   .S-7.
	   .|.|.
	   .L-J.
	   .....`

	expected := 0
	day := new(aoc.Day10)
	util.TestPart1(t, day, input, expected)
}
