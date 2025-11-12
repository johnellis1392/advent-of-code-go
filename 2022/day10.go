package aoc2022

import (
	"strconv"
	"strings"
)

type Day10 struct {
	input string
}

func (d *Day10) Year() string {
	return "2022"
}

func (d *Day10) Day() string {
	return "10"
}

func (d *Day10) Parse(input string) error {
	d.input = input
	return nil
}

func (d *Day10) Part1() any {
	sum := 0
	clock := 0
	cycleCheck := 20
	x := 1
	for _, line := range strings.Split(d.input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if line == "noop" {
			clock++
			if clock >= cycleCheck {
				cycleCheck += 40
				sum += x * clock
			}
		} else {
			n, _ := strconv.Atoi(strings.Split(line, " ")[1])
			clock++
			if clock >= cycleCheck {
				cycleCheck += 40
				sum += x * clock
			}
			clock++
			if clock >= cycleCheck {
				cycleCheck += 40
				sum += x * clock
			}
			x += n
		}
	}
	return sum
}

func tick(clock *int, cycleCheck *int, x int) string {
	pixel := *clock % 40
	*clock++
	res := ""

	if x-1 <= pixel && pixel <= x+1 {
		res += "#"
	} else {
		res += "."
	}

	if *clock == *cycleCheck {
		*cycleCheck += 40
		res += "\n"
	}

	return res
}

func (d *Day10) Part2() any {
	res := ""
	clock := 0
	cycleCheck := 40
	x := 1

	for _, line := range strings.Split(d.input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if line == "noop" {
			res += tick(&clock, &cycleCheck, x)
		} else {
			n, _ := strconv.Atoi(strings.Split(line, " ")[1])
			res += tick(&clock, &cycleCheck, x)
			res += tick(&clock, &cycleCheck, x)
			x += n
		}
	}

	return res
}
