package aoc2023

import (
	"regexp"
	"strconv"
	"strings"
)

type Day06 struct {
	input     string
	times     []int
	distances []int
}

func (d *Day06) Year() string {
	return "2023"
}

func (d *Day06) Day() string {
	return "06"
}

func (d *Day06) Parse(input string) error {
	re := regexp.MustCompile("[0-9]+")
	lines := strings.Split(input, "\n")
	timeStrings := re.FindAllString(strings.TrimSpace(strings.TrimPrefix(lines[0], "Time:")), -1)
	distanceStrings := re.FindAllString(strings.TrimSpace(strings.TrimPrefix(lines[1], "Distance:")), -1)

	var times, distances []int

	for i := 0; i < len(timeStrings); i++ {
		time, _ := strconv.Atoi(timeStrings[i])
		distance, _ := strconv.Atoi(distanceStrings[i])
		times = append(times, time)
		distances = append(distances, distance)
	}

	d.times = times
	d.distances = distances

	return nil
}

func (d *Day06) Part1() any {
	times, distances := d.times, d.distances
	res := 1
	for i := 0; i < len(times); i++ {
		time, distance := times[i], distances[i]
		n := 0
		for t := 1; t < time; t++ {
			speed := t
			dt := time - t
			dd := dt * speed
			if dd > distance {
				n++
			}
		}
		res *= n
	}
	return res
}

func readInput2(input string) (int, int) {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile("[0-9]")
	timeStrings := re.FindAllString(strings.TrimSpace(strings.TrimPrefix(lines[0], "Time:")), -1)
	distanceStrings := re.FindAllString(strings.TrimSpace(strings.TrimPrefix(lines[1], "Distance:")), -1)

	timeString := ""
	for _, ts := range timeStrings {
		timeString += ts
	}

	distanceString := ""
	for _, ds := range distanceStrings {
		distanceString += ds
	}

	time, _ := strconv.Atoi(timeString)
	distance, _ := strconv.Atoi(distanceString)

	return time, distance
}

func (d *Day06) Part2() any {
	time, distance := readInput2(d.input)
	n := 0
	for t := 1; t < time; t++ {
		speed := t
		dt := time - t
		dd := dt * speed
		if dd > distance {
			n++
		}
	}
	return n
}
