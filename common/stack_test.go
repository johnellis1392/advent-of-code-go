package common_test

import (
	"testing"

	"github.com/johnellis1392/advent-of-code-go/common"
	tu "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestStack(t *testing.T) {
	t.Run("Stack should report data correctly", func(t *testing.T) {
		stack := common.NewStack()
		tu.AssertTrue(t, stack.IsEmpty())
		tu.AssertFalse(t, stack.IsNotEmpty())
		tu.AssertEquals(t, 0, stack.Size())

		stack.Push(1)
		tu.AssertFalse(t, stack.IsEmpty())
		tu.AssertTrue(t, stack.IsNotEmpty())
		tu.AssertEquals(t, 1, stack.Size())

		stack.Push(2)
		tu.AssertFalse(t, stack.IsEmpty())
		tu.AssertTrue(t, stack.IsNotEmpty())
		tu.AssertEquals(t, 2, stack.Size())

		stack.Push(3)
		tu.AssertFalse(t, stack.IsEmpty())
		tu.AssertTrue(t, stack.IsNotEmpty())
		tu.AssertEquals(t, 3, stack.Size())
	})

	t.Run("Stack should mutate data correctly", func(t *testing.T) {
		stack := common.NewStack()
		tu.AssertNil(t, stack.Peek())

		stack.Push("1")
		tu.AssertEquals(t, "1", stack.Peek())

		stack.Push("2")
		tu.AssertEquals(t, "2", stack.Peek())

		stack.Push("3")
		tu.AssertEquals(t, "3", stack.Peek())
		tu.AssertEquals(t, 3, stack.Size())

		tu.AssertEquals(t, "3", stack.Pop())
		tu.AssertEquals(t, 2, stack.Size())
		tu.AssertEquals(t, "2", stack.Peek())

		tu.AssertEquals(t, "2", stack.Pop())
		tu.AssertEquals(t, 1, stack.Size())
		tu.AssertEquals(t, "1", stack.Peek())

		tu.AssertEquals(t, "1", stack.Pop())
		tu.AssertEquals(t, 0, stack.Size())
		tu.AssertNil(t, stack.Peek())
		tu.AssertNil(t, stack.Pop())
	})
}
