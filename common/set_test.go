package common_test

import (
	"testing"

	"github.com/johnellis1392/advent-of-code-go/common"
	tu "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestSet(t *testing.T) {
	t.Run("Test reporting methods", func(t *testing.T) {
		set := common.NewSet()
		tu.AssertEquals(t, 0, set.Size())
		tu.AssertTrue(t, set.IsEmpty())
		tu.AssertFalse(t, set.IsNotEmpty())

		set.Add("1")
		tu.AssertEquals(t, 1, set.Size())
		tu.AssertFalse(t, set.IsEmpty())
		tu.AssertTrue(t, set.IsNotEmpty())

		set.Add("2")
		tu.AssertEquals(t, 2, set.Size())
		tu.AssertFalse(t, set.IsEmpty())
		tu.AssertTrue(t, set.IsNotEmpty())

		set.Add("3")
		tu.AssertEquals(t, 3, set.Size())
		tu.AssertFalse(t, set.IsEmpty())
		tu.AssertTrue(t, set.IsNotEmpty())

		tu.AssertTrue(t, set.Contains("1"))
		tu.AssertTrue(t, set.Contains("2"))
		tu.AssertTrue(t, set.Contains("3"))
		tu.AssertFalse(t, set.Contains("4"))
		tu.AssertFalse(t, set.Contains("5"))
		tu.AssertFalse(t, set.Contains("6"))
	})

	t.Run("Test mutation methods", func(t *testing.T) {
		set := common.NewSet()
		set.Add("1")
		set.Add("2")
		set.Add("3")

		tu.AssertTrue(t, set.Contains("3"))
		set.Remove("3")
		tu.AssertFalse(t, set.Contains("3"))
		tu.AssertEquals(t, 2, set.Size())

		tu.AssertTrue(t, set.Contains("2"))
		set.Remove("2")
		tu.AssertFalse(t, set.Contains("2"))
		tu.AssertEquals(t, 1, set.Size())

		tu.AssertTrue(t, set.Contains("1"))
		set.Remove("1")
		tu.AssertFalse(t, set.Contains("1"))
		tu.AssertEquals(t, 0, set.Size())
	})

	t.Run("Sets should deduplicate elements", func(t *testing.T) {
		set := common.NewSet()
		set.Add("1")
		set.Add("2")
		set.Add("3")
		tu.AssertEquals(t, 3, set.Size())

		set.Add("1")
		tu.AssertEquals(t, 3, set.Size())

		set.Add("2")
		tu.AssertEquals(t, 3, set.Size())

		set.Add("3")
		tu.AssertEquals(t, 3, set.Size())
	})

	t.Run("Diff should calculate the diff of two sets", func(t *testing.T) {
		s1 := common.NewSetFrom([]any{"1", "2", "3"})
		s2 := common.NewSetFrom([]any{"3", "4", "5"})
		res := s1.Diff(s2)
		exp := common.NewSetFrom([]any{"1", "2"})
		tu.AssertEquals(t, exp, res)

		res = s2.Diff(s1)
		exp = common.NewSetFrom([]any{"4", "5"})
		tu.AssertEquals(t, exp, res)
	})

	t.Run("Union should calculate the union of two sets", func(t *testing.T) {
		s1 := common.NewSetFrom([]any{"1", "2", "3"})
		s2 := common.NewSetFrom([]any{"3", "4", "5"})
		res := s1.Union(s2)
		exp := common.NewSetFrom([]any{"1", "2", "3", "4", "5"})
		tu.AssertEquals(t, exp, res)
	})

	t.Run("Intersect should calculate the intersection of two sets", func(t *testing.T) {
		s1 := common.NewSetFrom([]any{"1", "2", "3"})
		s2 := common.NewSetFrom([]any{"3", "4", "5"})
		res := s1.Intersect(s2)
		exp := common.NewSetFrom([]any{"3"})
		tu.AssertEquals(t, exp, res)
	})

	t.Run("Overlaps should determine if two sets have any shared elements", func(t *testing.T) {
		s1 := common.NewSetFrom([]any{"1", "2", "3"})
		s2 := common.NewSetFrom([]any{"3", "4", "5"})
		tu.AssertTrue(t, s1.Overlaps(s2))

		s1.Remove("3")
		tu.AssertFalse(t, s1.Overlaps(s2))
	})

	t.Run("Xor should calculate the exclusive or of two sets", func(t *testing.T) {
		s1 := common.NewSetFrom([]any{"1", "2", "3"})
		s2 := common.NewSetFrom([]any{"3", "4", "5"})
		res := s1.Xor(s2)
		exp := common.NewSetFrom([]any{"1", "2", "4", "5"})
		tu.AssertEquals(t, exp, res)
	})
}
