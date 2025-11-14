package common

type Stack struct {
	data []any
}

func NewStack() *Stack {
	return &Stack{[]any{}}
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s *Stack) Push(v any) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() any {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[s.Size()-1]
	s.data = s.data[:s.Size()-1]
	return v
}

func (s *Stack) Peek() any {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.Size()-1]
}

func (s *Stack) Contains(v any) bool {
	for _, d := range s.data {
		if d == v {
			return true
		}
	}
	return false
}
