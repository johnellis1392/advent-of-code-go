package testutils

// Adds some missing assertions

import (
	"reflect"
	"slices"
	"testing"

	common "github.com/johnellis1392/advent-of-code-go/common"
)

func AssertEquals(t *testing.T, expected, actual any) {
	if expected == nil {
		if actual != nil {
			t.Errorf("Expected nil, but found %v", actual)
		}
	} else if e, ok := expected.(common.Base); ok {
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
