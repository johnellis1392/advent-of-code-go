package common

import (
	"fmt"
	"strings"
)

type Queue struct {
	data []any
}

func NewQueue() *Queue {
	return &Queue{[]any{}}
}

func NewQueueFrom(data []any) *Queue {
	return &Queue{data}
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) IsNotEmpty() bool {
	return !q.IsEmpty()
}

func (q *Queue) Peek() any {
	if q.IsEmpty() {
		return nil
	}
	return q.data[0]
}

func (q *Queue) Enqueue(v any) {
	q.data = append(q.data, v)
}

func (q *Queue) Dequeue() any {
	if q.IsEmpty() {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *Queue) Equals(o any) bool {
	if qo, ok := o.(*Queue); ok {
		if q.Size() != qo.Size() {
			return false
		}
		for i := 0; i < len(q.data); i++ {
			if q.data[i] != qo.data[i] {
				return false
			}
		}
		return true
	}
	return false
}

func (q *Queue) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	comma := false
	for _, v := range q.data {
		if comma {
			sb.WriteString(", ")
		} else {
			comma = true
		}
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	sb.WriteString("]")
	return sb.String()
}
