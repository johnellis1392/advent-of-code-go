package aoc2022

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day11 struct {
	input []Monkey
}

func (d *Day11) Year() string {
	return "2022"
}

func (d *Day11) Day() string {
	return "11"
}

func (d *Day11) Parse(input string) error {
	lines := strings.Split(input, "\n")
	var res []Monkey
	for i := 0; i < len(lines); i++ {
		monkeyLine := strings.TrimSpace(lines[i])
		if len(monkeyLine) == 0 {
			break
		}
		monkeyLine = strings.TrimPrefix(monkeyLine, "Monkey ")
		monkeyLine = monkeyLine[0 : len(monkeyLine)-1]
		monkeyId, _ := strconv.Atoi(monkeyLine)

		var items []uint64
		itemLine := strings.TrimSpace(lines[i+1])
		itemLine = strings.TrimPrefix(itemLine, "Starting items: ")
		itemStrings := strings.Split(itemLine, ", ")
		for _, s := range itemStrings {
			item, _ := strconv.Atoi(s)
			items = append(items, uint64(item))
		}

		opLine := strings.TrimSpace(lines[i+2])
		opLine = strings.TrimPrefix(opLine, "Operation: new = old ")
		var op Op
		n := opLine[2:]
		if opLine[0] == '+' {
			op = PlusOp{n}
		} else {
			op = MulOp{n}
		}

		testLine := strings.TrimSpace(lines[i+3])
		testLine = strings.TrimPrefix(testLine, "Test: divisible by ")
		test, _ := strconv.Atoi(testLine)

		trueLine := strings.TrimSpace(lines[i+4])
		trueLine = strings.TrimPrefix(trueLine, "If true: throw to monkey ")
		trueMonkey, _ := strconv.Atoi(trueLine)

		falseLine := strings.TrimSpace(lines[i+5])
		falseLine = strings.TrimPrefix(falseLine, "If false: throw to monkey ")
		falseMonkey, _ := strconv.Atoi(falseLine)

		res = append(res, Monkey{
			id:             monkeyId,
			items:          items,
			operation:      op,
			test:           uint64(test),
			trueMonkey:     trueMonkey,
			falseMonkey:    falseMonkey,
			numInspections: 0,
		})

		i += 6
	}

	d.input = res
	return nil
}

type Op interface {
	Eval(n uint64) uint64
	Value() string
	String() string
}

type PlusOp struct {
	i string
}

func (op PlusOp) Eval(n uint64) uint64 {
	if op.i == "old" {
		return n + n
	}
	v, _ := strconv.Atoi(op.i)
	return uint64(v) + n
}

func (op PlusOp) Value() string {
	return op.i
}

func (op PlusOp) String() string {
	return "+"
}

type MulOp struct {
	i string
}

func (op MulOp) Eval(n uint64) uint64 {
	if op.i == "old" {
		return n * n
	}
	v, _ := strconv.Atoi(op.i)
	return uint64(v) * n
}

func (op MulOp) Value() string {
	return op.i
}

func (op MulOp) String() string {
	return "*"
}

type Monkey struct {
	id                      int
	items                   []uint64
	operation               Op
	test                    uint64
	trueMonkey, falseMonkey int
	numInspections          uint64
}

func (m Monkey) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Monkey %d:\n", m.id))
	sb.WriteString(" Starting Items: ")
	for i, item := range m.items {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprint(item))
	}
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf(" Operation: new = old %s %s\n", m.operation.String(), m.operation.Value()))
	sb.WriteString(fmt.Sprintf(" Test: divisible by %d\n", m.test))
	sb.WriteString(fmt.Sprintf("  If true: throw to monkey %d\n", m.trueMonkey))
	sb.WriteString(fmt.Sprintf("  If false: throw to monkey %d\n", m.falseMonkey))
	return sb.String()
}

func (m *Monkey) Inspect(reliefEnabled bool, divisor uint64) (uint64, bool) {
	if len(m.items) == 0 {
		panic("Monkey has no items to inspect")
	}

	worryLevel := m.items[0]
	m.items = m.items[1:]

	worryLevel = m.operation.Eval(worryLevel)
	if reliefEnabled {
		worryLevel /= 3
	}
	m.numInspections++
	worryLevel = worryLevel % divisor
	return worryLevel, worryLevel%m.test == 0
}

func (m *Monkey) Receive(v uint64) {
	m.items = append(m.items, v)
}

// func dump(monkies []Monkey) {
// 	for _, monkey := range monkies {
// 		fmt.Println(monkey.String())
// 		fmt.Println()
// 	}
// }

func (d *Day11) Part1() any {
	monkies := d.input
	var divisor uint64 = 1
	for _, m := range monkies {
		divisor *= m.test
	}
	for round := 0; round < 20; round++ {
		for i := 0; i < len(monkies); i++ {
			monkey := &monkies[i]
			for len(monkey.items) > 0 {
				v, ok := monkey.Inspect(true, divisor)
				var target *Monkey
				if ok {
					target = &monkies[monkey.trueMonkey]
				} else {
					target = &monkies[monkey.falseMonkey]
				}
				target.Receive(v)
			}
		}
	}

	var inspections []uint64
	for _, monkey := range monkies {
		inspections = append(inspections, monkey.numInspections)
	}
	sort.Slice(inspections, func(i, j int) bool { return inspections[i] > inspections[j] })

	return inspections[0] * inspections[1]
}

func (d *Day11) Part2() any {
	monkies := d.input
	var divisor uint64 = 1
	for _, m := range monkies {
		divisor *= m.test
	}
	for round := 1; round <= 10000; round++ {
		for i := 0; i < len(monkies); i++ {
			monkey := &monkies[i]
			for len(monkey.items) > 0 {
				v, ok := monkey.Inspect(false, divisor)
				var target *Monkey
				if ok {
					target = &monkies[monkey.trueMonkey]
				} else {
					target = &monkies[monkey.falseMonkey]
				}
				target.Receive(v)
			}
		}
	}

	var inspections []uint64
	for _, monkey := range monkies {
		inspections = append(inspections, monkey.numInspections)
	}
	sort.Slice(inspections, func(i, j int) bool { return inspections[i] > inspections[j] })

	return inspections[0] * inspections[1]
}
