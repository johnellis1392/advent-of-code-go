package aoc2023

import (
	"strconv"
	"strings"
)

type Day02 struct {
	input []Game
}

func (d *Day02) Year() string {
	return "2023"
}

func (d *Day02) Day() string {
	return "02"
}

func (d *Day02) Parse(input string) error {
	var games []Game
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		var game Game

		parts := strings.Split(line, ":")
		id, _ := strings.CutPrefix(parts[0], "Game ")
		idn, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		game.id = idn

		var rolls [][]struct {
			n     int
			color string
		}

		for _, i := range strings.Split(parts[1], ";") {
			i = strings.TrimSpace(i)
			var rs []struct {
				n     int
				color string
			}

			for _, j := range strings.Split(i, ",") {
				j = strings.TrimSpace(j)
				roll := strings.Split(j, " ")
				n, err := strconv.Atoi(roll[0])
				if err != nil {
					panic(err)
				}
				color := roll[1]
				rs = append(rs, struct {
					n     int
					color string
				}{
					n:     n,
					color: color,
				})
			}
			rolls = append(rolls, rs)
		}

		game.rolls = rolls
		games = append(games, game)
	}
	d.input = games

	return nil
}

type Game struct {
	id    int
	rolls [][]struct {
		n     int
		color string
	}
}

func (d *Day02) Part1() any {
	games := d.input

	total_reds := 12
	total_greens := 13
	total_blues := 14

	sum := 0

outer:
	for _, game := range games {
		for _, round := range game.rolls {
			for _, roll := range round {
				switch roll.color {
				case "red":
					if roll.n > total_reds {
						continue outer
					}
				case "green":
					if roll.n > total_greens {
						continue outer
					}
				case "blue":
					if roll.n > total_blues {
						continue outer
					}
				}
			}
		}
		sum += game.id
	}

	return sum
}

func (d *Day02) Part2() any {
	games := d.input
	sum := 0

	for _, game := range games {
		reds, greens, blues := 0, 0, 0
		for _, round := range game.rolls {
			for _, roll := range round {
				switch roll.color {
				case "red":
					reds = max(reds, roll.n)
				case "green":
					greens = max(greens, roll.n)
				case "blue":
					blues = max(blues, roll.n)
				}
			}
		}
		sum += reds * greens * blues
	}

	return sum
}
