package common

type Day interface {
	Year() string
	Day() string
	Parse(input string) error
	Part1() any
	Part2() any
}
