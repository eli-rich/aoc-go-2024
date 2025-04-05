package day3

import (
	_ "embed"
	"regexp"
	"strconv"
)

//go:embed test.txt
var input string

var r *regexp.Regexp = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

type Solver struct{}

func (s Solver) Part1() (string, error) {
	matches := r.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		x, err := strconv.Atoi(match[1])
		if err != nil {
			return "", err
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			return "", err
		}
		sum += x * y
	}
	return strconv.Itoa(sum), nil
}

func (s Solver) Part2() (string, error) {

	return "", nil
}
