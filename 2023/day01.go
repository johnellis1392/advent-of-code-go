package aoc2023

import (
	"fmt"
	"regexp"
	"strings"
)

type Day01 struct {
	input string
}

func (d *Day01) Year() string {
	return "2023"
}

func (d *Day01) Day() string {
	return "01"
}

func (d *Day01) Parse(input string) error {
	d.input = input
	return nil
}

func (d *Day01) Part1() any {
	sum := 0
	re := regexp.MustCompile(`[^0-9]+`)

	for _, line := range strings.Split(d.input, "\n") {
		numbers := re.ReplaceAllString(line, "")
		if len(numbers) == 0 {
			continue
		}
		i := (int(numbers[0])-'0')*10 + (int(numbers[len(numbers)-1]) - '0')
		sum += i
	}

	return sum
}

func replace(input string) string {
	s := ""
	words := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4,
		"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	for i := 0; i < len(input); i++ {
		if '1' <= input[i] && input[i] <= '9' {
			s += fmt.Sprint(int(input[i] - '0'))
		}

		for k, v := range words {
			if strings.HasPrefix(input[i:], k) {
				s += fmt.Sprint(v)
				break
			}
		}
	}

	return s
}

func (d *Day01) Part2() any {
	sum := 0
	for _, line := range strings.Split(d.input, "\n") {
		if len(line) == 0 {
			continue
		}
		line = strings.TrimSpace(line)
		numbers := replace(line)
		i := (int(numbers[0])-'0')*10 + (int(numbers[len(numbers)-1]) - '0')
		sum += i
	}

	return sum
}
