package aoc2022

import (
	"strconv"
	"strings"
)

type Day05 struct {
	input string
}

func (d *Day05) Year() string {
	return "2022"
}

func (d *Day05) Day() string {
	return "05"
}

func (d *Day05) Parse(input string) error {
	d.input = input
	return nil
}

type stack struct {
	values []string
}

func newStack() *stack {
	s := stack{}
	return &s
}

func (s *stack) push(v string) {
	s.values = append(s.values, v)
}

func (s *stack) pop() *string {
	if len(s.values) == 0 {
		return nil
	}
	n := len(s.values)
	v := s.values[n-1]
	s.values = s.values[0 : n-1]
	return &v
}

func (s *stack) peek() *string {
	if len(s.values) == 0 {
		return nil
	}
	return &s.values[len(s.values)-1]
}

func (d *Day05) Part1() any {
	lines := strings.Split(d.input, "\n")
	n := 0
	for len(lines[n]) != 0 {
		n++
	}

	nStacks := (len(lines[n-1]) + 1) / 4
	stacks := make([]stack, nStacks)
	for i := 0; i < nStacks; i++ {
		stacks[i] = *newStack()
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < nStacks; j++ {
			c := string(lines[i][j*4+1])
			if c != " " {
				stacks[j].push(c)
			}
		}
	}

	for i := n + 1; i < len(lines); i++ {
		move := strings.Split(lines[i], " ")
		n, _ := strconv.Atoi(move[1])
		from, _ := strconv.Atoi(move[3])
		to, _ := strconv.Atoi(move[5])
		from -= 1
		to -= 1

		for j := 0; j < n; j++ {
			stacks[to].push(*stacks[from].pop())
		}
	}

	res := ""
	for _, stack := range stacks {
		res += *stack.peek()
	}
	return res
}

func (d *Day05) Part2() any {
	lines := strings.Split(d.input, "\n")
	n := 0
	for len(lines[n]) != 0 {
		n++
	}

	nStacks := (len(lines[n-1]) + 1) / 4
	stacks := make([]stack, nStacks)
	for i := 0; i < nStacks; i++ {
		stacks[i] = *newStack()
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < nStacks; j++ {
			c := string(lines[i][j*4+1])
			if c != " " {
				stacks[j].push(c)
			}
		}
	}

	temp := *newStack()
	for i := n + 1; i < len(lines); i++ {
		move := strings.Split(lines[i], " ")
		n, _ := strconv.Atoi(move[1])
		from, _ := strconv.Atoi(move[3])
		to, _ := strconv.Atoi(move[5])
		from -= 1
		to -= 1

		for j := 0; j < n; j++ {
			temp.push(*stacks[from].pop())
		}
		for j := 0; j < n; j++ {
			stacks[to].push(*temp.pop())
		}
	}

	res := ""
	for _, stack := range stacks {
		res += *stack.peek()
	}
	return res
}
