package aoc2023

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

type Day04 struct {
	input []Card
}

func (d *Day04) Year() string {
	return "2023"
}

func (d *Day04) Day() string {
	return "04"
}

func (d *Day04) Parse(input string) error {
	var res []Card
	re := regexp.MustCompile("[0-9]+")

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		var parts []string
		parts = strings.Split(line, ":")
		id := strings.TrimPrefix(parts[0], "Card ")

		parts = strings.Split(parts[1], " | ")
		winning := re.FindAllString(parts[0], -1)
		holding := re.FindAllString(parts[1], -1)

		res = append(res, Card{
			id:      id,
			winning: winning,
			holding: holding,
		})
	}

	d.input = res

	return nil
}

type Card struct {
	id      string
	winning []string
	holding []string
}

func (c *Card) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("Card %s: ", c.id))

	for i, n := range c.winning {
		if i > 0 {
			s.WriteString(" ")
		}
		s.WriteString(n)
	}

	s.WriteString(" | ")

	for i, n := range c.holding {
		if i > 0 {
			s.WriteString(" ")
		}
		s.WriteString(n)
	}

	return s.String()
}

func (card *Card) Score() int {
	n := 0
	for _, w := range card.winning {
		for _, h := range card.holding {
			if w == h {
				n++
			}
		}
	}
	return n
}

func dumpCards(cards []Card) {
	for _, card := range cards {
		fmt.Println(card.String())
	}
}

func (d *Day04) Part1() any {
	cards := d.input
	res := 0

	for _, card := range cards {
		n := card.Score()
		if n != 0 {
			res += int(math.Pow(2, float64(n-1)))
		}
	}

	return res
}

func dumpTotals(cardTotals []int) {
	for i, total := range cardTotals {
		fmt.Printf("Card %d: %d\n", i+1, total)
	}
}

func (d *Day04) Part2() any {
	cards := d.input
	cardTotals := make([]int, len(cards))
	for i := 0; i < len(cardTotals); i++ {
		cardTotals[i] = 1
	}

	for i, card := range cards {
		n := card.Score()
		instances := cardTotals[i]
		for j := i + 1; j < i+n+1; j++ {
			cardTotals[j] += instances
		}
	}

	res := 0
	for _, total := range cardTotals {
		res += total
	}

	return res
}
