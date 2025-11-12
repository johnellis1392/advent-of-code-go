package aoc2022

import (
	"strconv"
	"strings"
)

type Day04 struct {
	input string
}

func (d *Day04) Year() string {
	return "2022"
}

func (d *Day04) Day() string {
	return "04"
}

func (d *Day04) Parse(input string) error {
	d.input = input
	return nil
}

func (d *Day04) Part1() any {
	res := 0
	for _, line := range strings.Split(d.input, "\n") {
		line = strings.TrimSpace(line)
		assignments := strings.Split(line, ",")
		a1, a2 := strings.Split(assignments[0], "-"), strings.Split(assignments[1], "-")
		x1, _ := strconv.Atoi(a1[0])
		y1, _ := strconv.Atoi(a1[1])
		x2, _ := strconv.Atoi(a2[0])
		y2, _ := strconv.Atoi(a2[1])
		if x1 <= x2 && y2 <= y1 {
			res++
		} else if x2 <= x1 && y1 <= y2 {
			res++
		}
	}
	return res
}

func (d *Day04) Part2() any {
	res := 0
	for _, line := range strings.Split(d.input, "\n") {
		line = strings.TrimSpace(line)
		as := strings.Split(line, ",")
		p1, p2 := strings.Split(as[0], "-"), strings.Split(as[1], "-")
		x1, _ := strconv.Atoi(p1[0])
		y1, _ := strconv.Atoi(p1[1])
		x2, _ := strconv.Atoi(p2[0])
		y2, _ := strconv.Atoi(p2[1])
		if x1 <= x2 && x2 <= y1 || x1 <= y2 && y2 <= y1 {
			res += 1
		} else if x2 <= x1 && x1 <= y2 || x2 <= y1 && y1 <= y2 {
			res += 1
		}
	}
	return res
}
