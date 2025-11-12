package aoc2022

import (
	"strings"
)

func priority(c byte) int {
	if 'a' <= c && c <= 'z' {
		return int(c - 'a' + 1)
	} else {
		return int(c - 'A' + 27)
	}
}

type Day03 struct {
	input string
}

func (d *Day03) Year() string {
	return "2022"
}

func (d *Day03) Day() string {
	return "03"
}

func (d *Day03) Parse(input string) error {
	d.input = strings.TrimSpace(input)
	return nil
}

func (d *Day03) Part1() any {
	sum := 0
	for _, line := range strings.Split(d.input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		pivot := len(line) / 2
		var c byte
		left, right := line[0:pivot], line[len(line)-pivot:]

	loop:
		for l := 0; l < len(left); l++ {
			for r := 0; r < len(right); r++ {
				if left[l] == right[r] {
					c = left[l]
					break loop
				}
			}
		}

		sum += priority(c)
	}

	return sum
}

func intersection(s1, s2 string) string {
	res := ""
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				res += string(s1[i])
				break
			}
		}
	}
	return res
}

func (d *Day03) Part2() any {
	lines := strings.Split(d.input, "\n")
	sum := 0

	for i := 0; i < len(lines); i += 3 {
		s := strings.TrimSpace(lines[i])
		s = intersection(s, strings.TrimSpace(lines[i+1]))
		s = intersection(s, strings.TrimSpace(lines[i+2]))
		sum += priority(s[0])
	}

	return sum
}
