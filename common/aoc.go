package common

import (
	"fmt"
	"strconv"
)

type Day interface {
	Year() string
	Day() string
	Parse(input string) error
	Part1() any
	Part2() any
}

type SolutionsMap = map[string]Day

type SolutionsByYear = map[string]SolutionsMap

func SolutionsForYear(s SolutionsByYear, year string) SolutionsMap {
	if y, ok := s[year]; ok {
		return y
	}
	return nil
}

func SolutionForDay(s SolutionsByYear, year, day string) Day {
	if y, ok := s[year]; ok {
		if d, ok := y[day]; ok {
			return d
		}
	}
	return nil
}

func IsValidYear(year string) bool {
	if y, err := strconv.Atoi(year); err != nil {
		return false
	} else {
		return 2015 <= y && y <= 2024
	}
}

func IsValidDay(day string) bool {
	if d, err := strconv.Atoi(day); err != nil {
		return false
	} else {
		return 1 <= d && d <= 25
	}
}

func AllYears() []string {
	n := 2024 - 2015 + 1
	res := make([]string, n)
	for i := 0; i <= n; i++ {
		res[i] = fmt.Sprintf("%d", 2015+i)
	}
	return res
}

func AllDays() []string {
	res := make([]string, 25)
	for day := 1; day <= 25; day++ {
		res[day] = fmt.Sprintf("%02d", day)
	}
	return res
}
