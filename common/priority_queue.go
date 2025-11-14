package common

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type PriorityQueue struct {
	comparator func(any, any) bool
	data       []any
}

func NewPriorityQueue(comparator func(any, any) bool) *PriorityQueue {
	return &PriorityQueue{
		comparator: comparator,
		data:       []any{},
	}
}

func (p *PriorityQueue) Size() int {
	return len(p.data)
}

func (p *PriorityQueue) IsEmpty() bool {
	return p.Size() == 0
}

func (p *PriorityQueue) IsNotEmpty() bool {
	return !p.IsEmpty()
}

func (p *PriorityQueue) Enqueue(v any) {
	i := sort.Search(len(p.data), func(i int) bool {
		return p.comparator(v, p.data[i])
	})
	p.data = slices.Insert(p.data, i, v)
}

func (p *PriorityQueue) Dequeue() any {
	if p.IsEmpty() {
		return nil
	}
	v := p.data[0]
	p.data = p.data[1:]
	return v
}

func (p *PriorityQueue) Peek() any {
	if p.IsEmpty() {
		return nil
	}
	return p.data[0]
}

func (p *PriorityQueue) Equals(o any) bool {
	if po, ok := o.(*PriorityQueue); ok {
		if p.Size() != po.Size() {
			return false
		}
		for i := 0; i < p.Size(); i++ {
			if po.data[i] != p.data[i] {
				return false
			}
		}
		return true
	}
	return false
}

func (p *PriorityQueue) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	comma := false
	for _, v := range p.data {
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
