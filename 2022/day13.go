package aoc2022

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day13 struct {
	input string
}

func (d *Day13) Year() string {
	return "2022"
}

func (d *Day13) Day() string {
	return "13"
}

func (d *Day13) Parse(input string) error {
	d.input = input
	return nil
}

type Tree interface {
	IsLeaf() bool
	String() string
}

type Branch struct {
	vs []Tree
}

func NewBranch() *Branch {
	return &Branch{[]Tree{}}
}

func (b *Branch) IsLeaf() bool {
	return false
}

func (b *Branch) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, v := range b.vs {
		if i != 0 {
			sb.WriteString(",")
		}
		sb.WriteString(v.String())
	}
	sb.WriteString("]")
	return sb.String()
}

func (b *Branch) Add(t Tree) {
	b.vs = append(b.vs, t)
}

type Leaf struct {
	v int
}

func NewLeaf(v int) *Leaf {
	return &Leaf{v}
}

func (l *Leaf) IsLeaf() bool {
	return true
}

func (l *Leaf) String() string {
	return fmt.Sprint(l.v)
}

type Pair struct {
	left, right Tree
}

type Stack struct {
	vs []Tree
}

func (s *Stack) Push(t Tree) {
	s.vs = append(s.vs, t)
}

func (s *Stack) Pop() Tree {
	t := s.vs[len(s.vs)-1]
	s.vs = s.vs[0 : len(s.vs)-1]
	return t
}

func (s *Stack) Peek() Tree {
	return s.vs[len(s.vs)-1]
}

func (s *Stack) Empty() bool {
	return len(s.vs) == 0
}

func NewStack() *Stack {
	return &Stack{[]Tree{}}
}

func IsDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func parse(input string) Tree {
	stack := NewStack()
	stack.Push(NewBranch())

	var sb strings.Builder
	for i := 1; i < len(input)-1; i++ {
		c := input[i]
		switch {
		case c == '[':
			stack.Push(NewBranch())
		case c == ']':
			b := stack.Pop().(*Branch)
			if sb.Len() > 0 {
				n, _ := strconv.Atoi(sb.String())
				b.Add(NewLeaf(n))
				sb.Reset()
			}
			parent := stack.Peek().(*Branch)
			parent.Add(b)
		case IsDigit(c):
			sb.WriteByte(c)
		case c == ',':
			if sb.Len() > 0 {
				n, _ := strconv.Atoi(sb.String())
				b := stack.Peek().(*Branch)
				b.Add(NewLeaf(n))
				sb.Reset()
			}
		}
	}

	b := stack.Pop().(*Branch)
	if sb.Len() > 0 {
		n, _ := strconv.Atoi(sb.String())
		b.Add(NewLeaf(n))
	}
	return b
}

func readPairs(input string) []Pair {
	var res []Pair
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 3 {
		if len(strings.TrimSpace(lines[i])) == 0 {
			continue
		}
		left := parse(strings.TrimSpace(lines[i]))
		right := parse(strings.TrimSpace(lines[i+1]))
		res = append(res, Pair{left, right})
	}
	return res
}

func readSequences(input string) []Tree {
	var res []Tree
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		res = append(res, parse(line))
	}
	return res
}

func compareValues(left, right int) int {
	switch {
	case left < right:
		return 1
	case left > right:
		return -1
	default:
		return 0
	}
}

func compare(left, right Tree) int {
	switch {
	case left.IsLeaf() && right.IsLeaf():
		return compareValues(left.(*Leaf).v, right.(*Leaf).v)
	case left.IsLeaf() && !right.IsLeaf():
		return compare(&Branch{[]Tree{left}}, right)
	case !left.IsLeaf() && right.IsLeaf():
		return compare(left, &Branch{[]Tree{right}})
	default:
		leftBranch, rightBranch := left.(*Branch), right.(*Branch)

		i := 0
		for ; i < len(leftBranch.vs) && i < len(rightBranch.vs); i++ {
			switch compare(leftBranch.vs[i], rightBranch.vs[i]) {
			case 1:
				return 1
			case -1:
				return -1
			case 0:
				continue
			}
		}

		if i < len(rightBranch.vs) {
			return 1
		} else if i < len(leftBranch.vs) {
			return -1
		} else {
			return 0
		}
	}
}

func (d *Day13) Part1() any {
	pairs := readPairs(d.input)
	sum := 0
	for i, pair := range pairs {
		if compare(pair.left, pair.right) >= 0 {
			sum += i + 1
		}
	}
	return sum
}

func (d *Day13) Part2() any {
	sequences := readSequences(d.input)
	key1 := parse("[[2]]")
	key2 := parse("[[6]]")
	sequences = append(sequences, key1)
	sequences = append(sequences, key2)
	sort.Slice(sequences, func(i, j int) bool {
		s1, s2 := sequences[i], sequences[j]
		return compare(s1, s2) >= 0
	})

	res := 1
	for i, sequence := range sequences {
		s := sequence.String()
		if s == "[[2]]" || s == "[[6]]" {
			res *= i + 1
		}
	}
	return res
}
