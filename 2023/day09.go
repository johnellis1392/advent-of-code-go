package aoc2023

import (
	"strconv"
	"strings"
)

type Day09 struct {
	input [][]int
}

func (d *Day09) Year() string {
	return "2023"
}

func (d *Day09) Day() string {
	return "09"
}

func (d *Day09) Parse(input string) error {
	var result [][]int
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		var ns []int
		for _, v := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(v)
			ns = append(ns, n)
		}
		result = append(result, ns)
	}
	d.input = result

	return nil
}

func zeroes(sequence []int) bool {
	for _, v := range sequence {
		if v != 0 {
			return false
		}
	}
	return true
}

func differences(sequence []int) []int {
	diffs := make([]int, len(sequence)-1)
	for i := 0; i < len(diffs); i++ {
		diffs[i] = sequence[i+1] - sequence[i]
	}
	return diffs
}

func last(sequence []int) int {
	return sequence[len(sequence)-1]
}

func predict(sequence []int) int {
	if zeroes(sequence) {
		return 0
	}
	diffs := differences(sequence)
	next := predict(diffs)
	return last(sequence) + next
}

func (d *Day09) Part1() any {
	sequences := d.input
	sum := 0
	for _, sequence := range sequences {
		sum += predict(sequence)
	}
	return sum
}

func predictBackwards(sequence []int) int {
	if zeroes(sequence) {
		return 0
	}
	diffs := differences(sequence)
	prev := predictBackwards(diffs)
	return sequence[0] - prev
}

func (d *Day09) Part2() any {
	sequences := d.input
	sum := 0
	for _, sequence := range sequences {
		sum += predictBackwards(sequence)
	}
	return sum
}
