package common_test

import (
	"testing"

	"github.com/johnellis1392/advent-of-code-go/common"
	tu "github.com/johnellis1392/advent-of-code-go/testutils"
)

func TestQueue(t *testing.T) {
	t.Run("Queue getter methods", func(t *testing.T) {
		q := common.NewQueue()
		tu.AssertEquals(t, 0, q.Size())
		tu.AssertTrue(t, q.IsEmpty())
		tu.AssertFalse(t, q.IsNotEmpty())

		q.Enqueue("1")
		tu.AssertEquals(t, "1", q.Peek())
		tu.AssertEquals(t, 1, q.Size())
		tu.AssertFalse(t, q.IsEmpty())
		tu.AssertTrue(t, q.IsNotEmpty())

		q.Enqueue("2")
		tu.AssertEquals(t, "1", q.Peek())
		tu.AssertEquals(t, 2, q.Size())
		tu.AssertFalse(t, q.IsEmpty())
		tu.AssertTrue(t, q.IsNotEmpty())

		q.Enqueue("3")
		tu.AssertEquals(t, "1", q.Peek())
		tu.AssertEquals(t, 3, q.Size())
		tu.AssertFalse(t, q.IsEmpty())
		tu.AssertTrue(t, q.IsNotEmpty())
	})

	t.Run("Queue mutation methods", func(t *testing.T) {
		q := common.NewQueue()
		q.Enqueue("1")
		q.Enqueue("2")
		q.Enqueue("3")
		tu.AssertEquals(t, 3, q.Size())

		tu.AssertEquals(t, "1", q.Dequeue())
		tu.AssertEquals(t, 2, q.Size())

		tu.AssertEquals(t, "2", q.Dequeue())
		tu.AssertEquals(t, 1, q.Size())

		tu.AssertEquals(t, "3", q.Dequeue())
		tu.AssertEquals(t, 0, q.Size())
		tu.AssertNil(t, q.Dequeue())
	})
}
