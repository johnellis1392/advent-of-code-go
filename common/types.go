package common

type Day interface {
	Year() string
	Day() string
	Parse(input string) error
	Part1() any
	Part2() any
}

type Comparable interface {
	Equals(o any) bool
}

type Stringify interface {
	String() string
}

type Base interface {
	Comparable
	Stringify
}
