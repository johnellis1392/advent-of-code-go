package common_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/johnellis1392/advent-of-code-go/common"
	tu "github.com/johnellis1392/advent-of-code-go/testutils"
)

type testDatum struct {
	id       string
	priority int
}

func (td testDatum) Equals(o any) bool {
	if to, ok := o.(testDatum); ok {
		return td.id == to.id && td.priority == to.priority
	}
	return false
}

func (td testDatum) String() string {
	return fmt.Sprintf("{id=%v, priority=%d}", td.id, td.priority)
}

func TestPriorityQueue(t *testing.T) {

	testData := []testDatum{
		{"First", 1},
		{"Second", 2},
		{"Third", 3},
		{"Fourth", 4},
		{"Fifth", 5},
	}

	t.Run("PriorityQueue should sort by priority", func(t *testing.T) {
		vs := make([]testDatum, len(testData))
		copy(vs, testData)
		rand.Shuffle(len(vs), func(i, j int) {
			a := vs[i]
			vs[i] = vs[j]
			vs[j] = a
		})

		p := common.NewPriorityQueue(func(a, b any) bool {
			aa, _ := a.(testDatum)
			bb, _ := b.(testDatum)
			return aa.priority < bb.priority
		})

		tu.AssertEquals(t, 0, p.Size())
		tu.AssertTrue(t, p.IsEmpty())
		tu.AssertFalse(t, p.IsNotEmpty())

		p.Enqueue(vs[0])
		tu.AssertEquals(t, 1, p.Size())
		tu.AssertFalse(t, p.IsEmpty())
		tu.AssertTrue(t, p.IsNotEmpty())

		p.Enqueue(vs[1])
		tu.AssertEquals(t, 2, p.Size())
		tu.AssertFalse(t, p.IsEmpty())
		tu.AssertTrue(t, p.IsNotEmpty())

		p.Enqueue(vs[2])
		tu.AssertEquals(t, 3, p.Size())
		tu.AssertFalse(t, p.IsEmpty())
		tu.AssertTrue(t, p.IsNotEmpty())

		p.Enqueue(vs[3])
		tu.AssertEquals(t, 4, p.Size())
		tu.AssertFalse(t, p.IsEmpty())
		tu.AssertTrue(t, p.IsNotEmpty())

		p.Enqueue(vs[4])
		tu.AssertEquals(t, 5, p.Size())
		tu.AssertFalse(t, p.IsEmpty())
		tu.AssertTrue(t, p.IsNotEmpty())

		tu.AssertEquals(t, testData[0], p.Peek())
		tu.AssertEquals(t, testData[0], p.Dequeue())
		tu.AssertEquals(t, 4, p.Size())

		tu.AssertEquals(t, testData[1], p.Peek())
		tu.AssertEquals(t, testData[1], p.Dequeue())
		tu.AssertEquals(t, 3, p.Size())

		tu.AssertEquals(t, testData[2], p.Peek())
		tu.AssertEquals(t, testData[2], p.Dequeue())
		tu.AssertEquals(t, 2, p.Size())

		tu.AssertEquals(t, testData[3], p.Peek())
		tu.AssertEquals(t, testData[3], p.Dequeue())
		tu.AssertEquals(t, 1, p.Size())

		tu.AssertEquals(t, testData[4], p.Peek())
		tu.AssertEquals(t, testData[4], p.Dequeue())
		tu.AssertEquals(t, 0, p.Size())
		tu.AssertTrue(t, p.IsEmpty())
		tu.AssertFalse(t, p.IsNotEmpty())
		tu.AssertNil(t, p.Peek())
		tu.AssertNil(t, p.Dequeue())
	})
}
