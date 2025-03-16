package day1

import (
	_ "embed"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var lines = strings.Split(strings.TrimSpace(input), "\n")
var size = len(lines)

type Solver struct{}

func (s Solver) Part1() (string, error) {
	var left []int
	var right []int

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		leftVal, err := strconv.Atoi(tokens[0])
		if err != nil {
			return "Error: could not Atoi leftVal: " + tokens[0], err
		}
		last := len(tokens) - 1
		rightVal, err := strconv.Atoi(tokens[last])
		if err != nil {
			return "Error: could not Atoi rightVal: " + tokens[last], err
		}
		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i := range size {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	return strconv.Itoa(sum), nil
}

func (s Solver) Part2() (string, error) {
	var left []int
	var right []int

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		leftVal, err := strconv.Atoi(tokens[0])
		if err != nil {
			return "Error: could not Atoi leftVal: " + tokens[0], err
		}
		last := len(tokens) - 1
		rightVal, err := strconv.Atoi(tokens[last])
		if err != nil {
			return "Error: could not Atoi rightVal: " + tokens[last], err
		}
		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	similarity := make(map[int]int)

	leftUnique := slices.Clone(left)
	leftUnique = slices.Compact(leftUnique)

	for _, value := range leftUnique {
		// count occurances of value in right
		similarity[value] = 0
		for _, rightVal := range right {
			if rightVal == value {
				similarity[value]++
			}
		}
	}
	sum := 0
	for _, value := range left {
		sum += value * similarity[value]
	}

	return strconv.Itoa(sum), nil
}
