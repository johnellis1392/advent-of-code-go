package aoc2023

import (
	"sort"
	"strconv"
	"strings"
)

type Day07 struct {
	input []Hand
}

func (d *Day07) Year() string {
	return "2023"
}

func (d *Day07) Day() string {
	return "07"
}

func (d *Day07) Parse(input string) error {
	var hands []Hand
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		hands = append(hands, NewHand(line))
	}
	d.input = hands

	return nil
}

type Hand struct {
	cards    string
	bid      int
	values   []int
	handType int
}

func (h Hand) String() string {
	return h.cards + " bid: " + strconv.Itoa(h.bid)
}

func (h Hand) Less(other Hand) bool {
	if h.handType < other.handType {
		return true
	} else if h.handType > other.handType {
		return false
	}
	for i := 0; i < len(h.values); i++ {
		if h.values[i] != other.values[i] {
			return h.values[i] < other.values[i]
		}
	}
	return false
}

const FIVE_OF_A_KIND = 6
const FOUR_OF_A_KIND = 5
const FULL_HOUSE = 4
const THREE_OF_A_KIND = 3
const TWO_PAIR = 2
const ONE_PAIR = 1
const HIGH_CARD = 0

func CalcHandType(cards []int) int {
	// Calculate number of like-cards
	groups := make(map[int]int)
	for i := 0; i < len(cards); i++ {
		if v, ok := groups[cards[i]]; ok {
			groups[cards[i]] = v + 1
		} else {
			groups[cards[i]] = 1
		}
	}

	// Get number of cards per value, and sort descending
	var values []int
	for _, v := range groups {
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })

	// Determine the type of hand by the number of cards
	switch {
	case values[0] == 5: // Five of a Kind
		return FIVE_OF_A_KIND
	case values[0] == 4: // Four of a Kind
		return FOUR_OF_A_KIND
	case values[0] == 3 && values[1] == 2: // Full House
		return FULL_HOUSE
	case values[0] == 3 && values[1] == 1: // Three of a Kind
		return THREE_OF_A_KIND
	case values[0] == 2 && values[1] == 2: // Two Pair
		return TWO_PAIR
	case values[0] == 2 && values[1] == 1: // One Pair
		return ONE_PAIR
	default: // High Card
		return HIGH_CARD
	}
}

func cardRank(c byte) int {
	switch c {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	// case 'J':
	// 	return 11
	case 'T':
		return 10
	case '9':
		return 9
	case '8':
		return 8
	case '7':
		return 7
	case '6':
		return 6
	case '5':
		return 5
	case '4':
		return 4
	case '3':
		return 3
	case '2':
		return 2
	case 'J':
		return 1
	default:
		return 0
	}
}

func NewHand(input string) Hand {
	parts := strings.Split(input, " ")
	bid, _ := strconv.Atoi(parts[1])
	var values []int
	for i := 0; i < len(parts[0]); i++ {
		values = append(values, cardRank(parts[0][i]))
	}
	return Hand{
		cards:    parts[0],
		bid:      bid,
		values:   values,
		handType: CalcHandType(values),
	}
}

func (d *Day07) Part1() any {
	hands := d.input
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Less(hands[j])
	})
	res := 0
	for i, h := range hands {
		res += (i + 1) * h.bid
	}
	return res
}

func (d *Day07) Part2() any {
	return 0
}
