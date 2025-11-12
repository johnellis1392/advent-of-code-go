package testutils

import (
	"testing"

	common "github.com/johnellis1392/advent-of-code-go/common"
)

func TestPart1(t *testing.T, day common.Day, input string, expected any) {
	day.Parse(input)
	actual := day.Part1()
	if actual != expected {
		t.Errorf("part1: expected %v, got %v", expected, actual)
	}
}

func TestPart2(t *testing.T, day common.Day, input string, expected any) {
	day.Parse(input)
	actual := day.Part2()
	if actual != expected {
		t.Errorf("part2: expected %v, got %v", expected, actual)
	}
}
