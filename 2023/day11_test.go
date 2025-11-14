package aoc2023_test

import (
	"fmt"
	"testing"

	aoc "github.com/johnellis1392/advent-of-code-go/2023"
	common "github.com/johnellis1392/advent-of-code-go/common"
	tu "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestDay11_Part1(t *testing.T) {
	input :=
		`...#......
	.......#..
	#.........
	..........
	......#...
	.#........
	.........#
	..........
	.......#..
	#...#.....`

	//     0123456789
	//   ------------
	// 0 | ...0......
	// 1 | .......1..
	// 2 | 2.........
	// 3 | ..........
	// 4 | ......3...
	// 5 | .4........
	// 6 | .........5
	// 7 | ..........
	// 8 | .......6..
	// 9 | 7...8.....
	t.Run("Test Parse", func(t *testing.T) {
		d := new(aoc.Day11)
		d.Parse(input)
		tu.AssertEquals(t, 9, len(d.Galaxies()))
		expectedGalaxies := []common.Point{
			{X: 3, Y: 0},
			{X: 7, Y: 1},
			{X: 0, Y: 2},
			{X: 6, Y: 4},
			{X: 1, Y: 5},
			{X: 9, Y: 6},
			{X: 7, Y: 8},
			{X: 0, Y: 9},
			{X: 4, Y: 9},
		}
		tu.AssertEquals(t, expectedGalaxies, d.Galaxies())
		tu.AssertEquals(t, []bool{false, false, false, true, false, false, false, true, false, false}, d.NullRows())
		tu.AssertEquals(t, []bool{false, false, true, false, false, true, false, false, true, false}, d.NullCols())
		tu.AssertEquals(t, 10, d.Width())
		tu.AssertEquals(t, 10, d.Height())
	})

	t.Run("PathLength should return correct path lengths", func(t *testing.T) {
		d := new(aoc.Day11)
		d.Parse(input)

		type testData struct {
			from, to int
			expected uint64
		}
		testCases := []testData{
			{0, 1, 5},
			{1, 2, 8},
			{2, 3, 8},
			{3, 4, 6},
			{4, 5, 9},
			{5, 6, 4},
			{6, 7, 8},
			{7, 8, 4},
		}
		for _, tc := range testCases {
			name := fmt.Sprintf("%d->%d = %d", tc.from, tc.to, tc.expected)
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				from, to, expected := tc.from, tc.to, tc.expected
				tu.AssertEquals(t, expected, d.PathLength(from, to, 1))
			})
		}
	})

	t.Run("PathLength should return path lengths with increments for null spaces", func(t *testing.T) {
		d := new(aoc.Day11)
		d.Parse(input)
		const nullSpaceSize = 2

		type testCase struct {
			from, to int
			expected uint64
		}

		// 5->9 = 9
		// 1->7 = 15
		// 3->6 = 17
		// 8->9 = 5
		testCases := []testCase{
			{4, 8, 9},
			{0, 6, 15},
			{2, 5, 17},
			{7, 8, 5},
		}

		for _, tc := range testCases {
			name := fmt.Sprintf("%d->%d = %d (nullSpaceSize=2)", tc.from, tc.to, tc.expected)
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				from, to, expected := tc.from, tc.to, tc.expected
				tu.AssertEquals(t, expected, d.PathLength(from, to, nullSpaceSize))
			})
		}
	})

	t.Run("Test Data 1", func(t *testing.T) {
		d := new(aoc.Day11)
		expected := uint64(374)
		tu.TestPart1(t, d, input, expected)
	})
}
