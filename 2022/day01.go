package aoc2022

import (
	"sort"
	"strconv"
	"strings"
)

type Day01 struct {
	input string
}

func (d *Day01) Year() string {
	return "2022"
}

func (d *Day01) Day() string {
	return "01"
}

func (d *Day01) Parse(input string) error {
	d.input = strings.TrimSpace(input)
	return nil
}

func (d *Day01) Part1() any {
	res := 0
	sum := 0
	for _, line := range strings.Split(d.input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			res = max(res, sum)
			sum = 0
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		sum += n
	}
	res = max(sum, res)
	return res
}

func (d *Day01) Part2() any {
	var sums []int
	sum := 0
	for _, line := range strings.Split(d.input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			sums = append(sums, sum)
			sum = 0
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		sum += n
	}

	sums = append(sums, sum)
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	sum = 0
	for i := 0; i < 3; i++ {
		sum += sums[i]
	}
	return sum
}
