package aoc2023

import (
	"fmt"
	"strings"
)

type Day08 struct {
	dir   Direction
	insts map[string]Inst
}

func (d *Day08) Year() string {
	return "2023"
}

func (d *Day08) Day() string {
	return "08"
}

func (d *Day08) Parse(input string) error {
	lines := strings.Split(input, "\n")
	dirLine := lines[0]
	dirLine = strings.TrimSpace(dirLine)

	insts := make(map[string]Inst)

	for i := 2; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " = ")
		id := parts[0]
		dirString := strings.TrimSpace(parts[1])
		dirString = strings.TrimPrefix(dirString, "(")
		dirString = strings.TrimSuffix(dirString, ")")
		dirs := strings.Split(dirString, ", ")
		left, right := dirs[0], dirs[1]

		insts[id] = Inst{id, left, right}
	}

	d.dir = NewDirection(dirLine)
	d.insts = insts

	return nil
}

type Direction struct {
	i    int
	dirs string
}

func NewDirection(s string) Direction {
	return Direction{0, s}
}

func (d *Direction) Reset() {
	d.i = 0
}

func (d *Direction) Next() string {
	dir := string(d.dirs[d.i])
	d.i = (d.i + 1) % len(d.dirs)
	return dir
}

type Inst struct {
	id    string
	left  string
	right string
}

func (inst Inst) String() string {
	return fmt.Sprintf("%s = (%s, %s)", inst.id, inst.left, inst.right)
}

func (inst Inst) IsStart() bool {
	return inst.id == "AAA"
}

func (inst Inst) IsEnd() bool {
	return inst.id == "ZZZ"
}

func (d *Day08) Part1() any {
	dirs, insts := d.dir, d.insts
	n := 0
	current := insts["AAA"]
	for current.id != "ZZZ" {
		nextDir := dirs.Next()
		if nextDir == "R" {
			current = insts[current.right]
		} else {
			current = insts[current.left]
		}
		n++
	}
	return n
}

func isEnd(insts []Inst) bool {
	for _, inst := range insts {
		if !inst.IsEnd() {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(vs []int) int {
	res := vs[0] * vs[1] / gcd(vs[0], vs[1])
	for i := 2; i < len(vs); i++ {
		res = res * vs[i] / gcd(res, vs[i])
	}
	return res
}

func (d *Day08) Part2() any {
	dirs, insts := d.dir, d.insts
	var currentInsts []Inst
	for _, inst := range insts {
		if inst.IsStart() {
			currentInsts = append(currentInsts, inst)
		}
	}

	// Rote solution, too inefficient to work
	// n := 0
	// for !isEnd(currentInsts) {
	// 	n++
	// 	dir := dirs.Next()
	// 	for i := 0; i < len(currentInsts); i++ {
	// 		inst := currentInsts[i]
	// 		if dir == "R" {
	// 			currentInsts[i] = insts[inst.right]
	// 		} else {
	// 			currentInsts[i] = insts[inst.left]
	// 		}
	// 	}
	// }

	// LCM solution; works like a charm
	endStates := make([]int, len(currentInsts))
	for i := 0; i < len(currentInsts); i++ {
		dirs.Reset()
		inst := currentInsts[i]
		for !inst.IsEnd() {
			dir := dirs.Next()
			endStates[i]++
			if dir == "R" {
				inst = insts[inst.right]
			} else {
				inst = insts[inst.left]
			}
		}
	}

	res := lcm(endStates)

	return res
}
