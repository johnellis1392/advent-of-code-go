package common

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
