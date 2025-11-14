package common

import (
	"fmt"
	"strings"
)

type Set struct {
	vs map[any]bool
}

func NewSet() *Set {
	return &Set{make(map[any]bool)}
}

func NewSetFrom(vs []any) *Set {
	m := make(map[any]bool)
	for _, v := range vs {
		m[v] = true
	}
	return &Set{m}
}

func (s *Set) Size() int {
	return len(s.vs)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s *Set) First() any {
	for p := range s.vs {
		return p
	}
	return nil
}

func (s *Set) Add(p any) {
	s.vs[p] = true
}

func (s *Set) Remove(p any) {
	delete(s.vs, p)
}

func (s *Set) Contains(p any) bool {
	_, ok := s.vs[p]
	return ok
}

func (s *Set) Diff(o *Set) (res *Set) {
	res = NewSet()
	s.ForEach(func(p any) {
		if !o.Contains(p) {
			res.Add(p)
		}
	})
	return res
}

func (s *Set) Union(o *Set) (res *Set) {
	res = NewSet()
	s.ForEach(func(p any) {
		res.Add(p)
	})
	o.ForEach(func(p any) {
		res.Add(p)
	})
	return res
}

func (s *Set) Xor(o *Set) (res *Set) {
	res = NewSet()
	s.ForEach(func(p any) {
		if !o.Contains(p) {
			res.Add(p)
		}
	})
	o.ForEach(func(p any) {
		if !s.Contains(p) {
			res.Add(p)
		}
	})
	return res
}

func (s *Set) Intersect(o *Set) (res *Set) {
	res = NewSet()
	s.ForEach(func(p any) {
		if o.Contains(p) {
			res.Add(p)
		}
	})
	return res
}

func (s *Set) Overlaps(o *Set) bool {
	return s.Intersect(o).Size() > 0
}

func (s *Set) Equals(o any) bool {
	if so, ok := o.(*Set); ok {
		if s.Size() != so.Size() {
			return false
		}
		for v := range s.vs {
			if !so.Contains(v) {
				return false
			}
		}
		return true
	}
	return false
}

func (s *Set) ForEach(f func(p any)) {
	for p := range s.vs {
		f(p)
	}
}

func (s *Set) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	comma := false
	for k, _ := range s.vs {
		if comma {
			sb.WriteString(", ")
		} else {
			comma = false
		}
		if kb, ok := k.(Base); ok {
			sb.WriteString(kb.String())
		} else {
			sb.WriteString(fmt.Sprintf("%v", k))
		}
	}
	sb.WriteString("}")
	return sb.String()
}
