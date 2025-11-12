package aoc2022

type Day06 struct {
	input string
}

func (d *Day06) Year() string {
	return "2022"
}

func (d *Day06) Day() string {
	return "06"
}

func (d *Day06) Parse(input string) error {
	d.input = input
	return nil
}

func (d *Day06) Part1() any {
outer:
	for i := 4; i < len(d.input); i++ {
		for j := i - 4; j < i-1; j++ {
			for k := j + 1; k < i; k++ {
				if d.input[j] == d.input[k] {
					continue outer
				}
			}
		}
		return i
	}
	return -1
}

func (d *Day06) Part2() any {
	const n = 14
outer:
	for i := n; i < len(d.input); i++ {
		for j := i - n; j < i-1; j++ {
			for k := j + 1; k < i; k++ {
				if d.input[j] == d.input[k] {
					continue outer
				}
			}
		}
		return i
	}
	return -1
}
