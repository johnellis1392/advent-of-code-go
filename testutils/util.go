package testutils

import (
	"reflect"
	"slices"
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

func AssertEquals(t *testing.T, expected, actual any) {
	if e, ok := expected.(common.Base); ok {
		if !e.Equals(actual) {
			t.Errorf("Expected %v, but found %v", e.String(), actual)
		}
	} else if reflect.TypeOf(expected).Kind() == reflect.Slice {
		e2, _ := expected.([]any)
		a2, _ := actual.([]any)
		if !slices.Equal(e2, a2) {
			t.Errorf("Expected %v, but found %v", e.String(), actual)
		}
	} else {
		if expected != actual {
			t.Errorf("Expected %v, but found %v", expected, actual)
		}
	}
}

func AssertTrue(t *testing.T, v bool) {
	AssertEquals(t, true, v)
}

func AssertFalse(t *testing.T, v bool) {
	AssertEquals(t, false, v)
}

func AssertNil(t *testing.T, v any) {
	AssertEquals(t, nil, v)
}
