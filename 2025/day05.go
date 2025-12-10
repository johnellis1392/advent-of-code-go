package aoc2025

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	from, to int64
}

func (r Range) Contains(v int64) bool {
	return r.from <= v && v <= r.to
}

func (r Range) ContainsRange(o Range) bool {
	return r.Contains(o.from) && r.Contains(o.to)
}

func (r Range) Before(o Range) bool {
	return r.to < o.from
}

func (r Range) After(o Range) bool {
	return o.to < r.from
}

func (r Range) Overlaps(o Range) bool {
	return r.Contains(o.from) || r.Contains(o.to)
}

func (r Range) ExtendsBefore(o Range) bool {
	return o.Contains(r.to)
}

func (r Range) ExtendsAfter(o Range) bool {
	return o.Contains(r.from)
}

func (r Range) AdjacentBefore(o Range) bool {
	return o.from == r.to+1
}

func (r Range) Adjacent(o Range) bool {
	return r.AdjacentBefore(o) || r.AdjacentAfter(o)
}

func (r Range) AdjacentAfter(o Range) bool {
	return r.from == o.to+1
}

func (r Range) Size() int64 {
	return r.to - r.from + 1
}

type Tree struct {
	root Node
}

func (t *Tree) Insert(r Range) {
	if t.root == nil {
		t.root = &Leaf{r}
	} else {
		t.root = t.root.Insert(r)
	}
}

func (t *Tree) Add(from, to int64) {
	t.Insert(Range{from, to})
}

func (t *Tree) Count() int64 {
	if t.root == nil {
		return 0
	}
	return t.root.Count()
}

type Node interface {
	Insert(r Range) Node
	Range() Range
	Count() int64
}

type Branch struct {
	left, right Node
}

type Leaf struct {
	r Range
}

func (n *Leaf) Range() Range {
	return n.r
}

func (n *Leaf) Insert(r Range) Node {
	switch {
	case r.Before(n.r):
		return &Branch{&Leaf{r}, n}
	case r.After(n.r):
		return &Branch{n, &Leaf{r}}
	case r.ContainsRange(n.r):
		n.r = r
		return n
	case n.r.ContainsRange(r):
		return n
	case r.Overlaps(n.r) || r.Adjacent(n.r):
		return &Leaf{Range{min(r.from, n.r.from), max(r.to, n.r.to)}}
	default:
		return n
	}
}

func (n *Leaf) Count() int64 {
	return n.r.Size()
}

func (node *Branch) Range() Range {
	return Range{node.left.Range().from, node.right.Range().to}
}

func (node *Branch) Insert(that Range) Node {
	leftRange := node.left.Range()
	rightRange := node.right.Range()
	switch {
	case that.ContainsRange(leftRange) && that.ContainsRange(rightRange):
		return &Leaf{that}
	case that.ExtendsAfter(leftRange) && that.ExtendsBefore(rightRange):
		return &Leaf{Range{leftRange.from, rightRange.to}}
	case that.AdjacentAfter(leftRange) && that.ExtendsBefore(rightRange):
		return &Leaf{Range{leftRange.from, rightRange.to}}
	case that.Before(leftRange):
		node.left = node.left.Insert(that)
	case that.After(rightRange):
		node.right = node.right.Insert(that)
	case that.Overlaps(leftRange):
		node.left = node.left.Insert(that)
	case that.Overlaps(rightRange):
		node.right = node.right.Insert(that)
	default:
		// Otherwise, we know the range is between leftRange and rightRange,
		// so choose whichever is closer.
		if that.from-leftRange.to < rightRange.from-that.to {
			node.left = node.left.Insert(that)
		} else {
			node.right = node.right.Insert(that)
		}
	}
	return node
}

func (node *Branch) Count() int64 {
	return node.left.Count() + node.right.Count()
}

type Day05 struct {
	input  string
	ranges []Range
	seeds  []int64
}

func (d *Day05) Year() string {
	return "2025"
}

func (d *Day05) Day() string {
	return "05"
}

func (d *Day05) Parse(input string) error {
	d.input = input
	parts := strings.Split(input, "\n\n")
	rangesString := parts[0]
	seedsString := parts[1]

	var ranges []Range
	for r := range strings.SplitSeq(rangesString, "\n") {
		rr := strings.Split(r, "-")
		from, _ := strconv.ParseInt(rr[0], 10, 64)
		to, _ := strconv.ParseInt(rr[1], 10, 64)
		ranges = append(ranges, Range{from, to})
	}

	var seeds []int64
	for s := range strings.SplitSeq(seedsString, "\n") {
		ss, _ := strconv.ParseInt(s, 10, 64)
		seeds = append(seeds, ss)
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return cmp.Compare(a.from, b.from)
	})
	d.ranges = ranges
	d.seeds = seeds
	return nil
}

func (d *Day05) Part1() any {
	res := 0
	for _, s := range d.seeds {
		for _, r := range d.ranges {
			if r.Contains(s) {
				res += 1
				break
			}
		}
	}
	return res
}

func (d *Day05) Part2() any {
	tree := Tree{nil}
	for _, r := range d.ranges {
		tree.Insert(r)
	}
	return tree.Count()
}
