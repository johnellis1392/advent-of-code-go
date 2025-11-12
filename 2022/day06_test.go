package aoc2022_test

import (
	"testing"

	aoc2022 "github.com/johnellis1392/advent-of-code-go/2022"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

var tests []string = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func TestDay06_Part1(t *testing.T) {
	input := tests[0]
	expected := 0
	day := new(aoc2022.Day06)
	util.TestPart1(t, day, input, expected)
}
