package aoc2022

import (
	"strings"
)

type Day02 struct {
	input string
}

func (d *Day02) Year() string {
	return "2022"
}

func (d *Day02) Day() string {
	return "02"
}

func (d *Day02) Parse(input string) error {
	d.input = strings.TrimSpace(input)
	return nil
}

func (d *Day02) Part1() any {
	win_score := 6
	draw_score := 3
	lose_score := 0
	rock_score := 1
	paper_score := 2
	scissor_score := 3

	sum := 0
	for _, line := range strings.Split(d.input, "\n") {
		line = strings.TrimSpace(line)
		switch line {
		case "A X":
			sum += draw_score + rock_score
		case "A Y":
			sum += win_score + paper_score
		case "A Z":
			sum += lose_score + scissor_score
		case "B X":
			sum += lose_score + rock_score
		case "B Y":
			sum += draw_score + paper_score
		case "B Z":
			sum += win_score + scissor_score
		case "C X":
			sum += win_score + rock_score
		case "C Y":
			sum += lose_score + paper_score
		case "C Z":
			sum += draw_score + scissor_score
		}
	}

	return sum
}

func (d *Day02) Part2() any {
	win_score := 6
	draw_score := 3
	lose_score := 0
	rock_score := 1
	paper_score := 2
	scissor_score := 3

	sum := 0

	for _, line := range strings.Split(d.input, "\n") {
		line = strings.TrimSpace(line)
		switch line {
		case "A X":
			sum += lose_score + scissor_score
		case "A Y":
			sum += draw_score + rock_score
		case "A Z":
			sum += win_score + paper_score
		case "B X":
			sum += lose_score + rock_score
		case "B Y":
			sum += draw_score + paper_score
		case "B Z":
			sum += win_score + scissor_score
		case "C X":
			sum += lose_score + paper_score
		case "C Y":
			sum += draw_score + scissor_score
		case "C Z":
			sum += win_score + rock_score
		}
	}

	return sum
}
