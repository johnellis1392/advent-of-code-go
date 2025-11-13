package aoc2022_test

import (
	"fmt"
	"testing"

	aoc2022 "github.com/johnellis1392/advent-of-code-go/2022"
	util "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay06_Part1(t *testing.T) {
	type testCase struct {
		input    string
		expected int
	}

	tests := []testCase{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}
	day := new(aoc2022.Day06)
	for i, testCase := range tests {
		testId := fmt.Sprintf("Part 1, test %d", i)
		t.Run(testId, func(t *testing.T) {
			util.TestPart1(t, day, testCase.input, testCase.expected)
		})
	}
}
