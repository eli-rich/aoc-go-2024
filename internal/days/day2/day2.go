package day2

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var lines = strings.Split(strings.TrimSpace(input), "\n")

func absInt(a, b int) int {
	result := a - b
	if result < 0 {
		return -result
	}
	return result
}

type Solver struct{}

func (s Solver) Part1() (string, error) {
	safeReports := 0
	for _, line := range lines {
		values := strings.Split(line, " ")
		var prev int
		first := true
		increasing := true
		decreasing := true
		safe := true
		for _, valueString := range values {
			value, err := strconv.Atoi(valueString)
			if err != nil {
				return "", err
			}
			if first {
				prev = value
				first = false
				continue
			}
			// check to ensure difference is 1-3 inclusive
			diff := absInt(prev, value)
			// fmt.Printf("prev: %d, value: %d, diff: %d\n", prev, value, diff)
			if diff < 1 || diff > 3 {
				safe = false
				break
			}
			if value < prev {
				increasing = false
			}
			if value > prev {
				decreasing = false
			}
			if !increasing && !decreasing {
				safe = false
				break
			}
			prev = value
		}
		if safe {
			safeReports++
		}

	}
	return strconv.Itoa(safeReports), nil
}

func (s Solver) Part2() (string, error) {
	safeReports := 0
	for _, line := range lines {
		numStrings := strings.Fields(line)
		nums := make([]int, len(numStrings))
		for i, s := range numStrings {
			nums[i], _ = strconv.Atoi(s)
		}
		if isSafe(nums) {
			safeReports++
		}
	}
	return strconv.Itoa(safeReports), nil
}

func isSafe(nums []int) bool {
	if isMonotonic(nums) && validDiffs(nums) {
		return true
	}
	for i := range nums {
		tmp := make([]int, 0, len(nums)-1)
		tmp = append(tmp, nums[:i]...)
		tmp = append(tmp, nums[i+1:]...)

		if isMonotonic(tmp) && validDiffs(tmp) {
			return true
		}
	}
	return false
}

func isMonotonic(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	increasing := nums[0] < nums[1]
	decreasing := nums[0] > nums[1]

	if !increasing && !decreasing {
		return false
	}

	for i := 2; i < len(nums); i++ {
		if increasing && nums[i] <= nums[i-1] {
			return false
		}
		if decreasing && nums[i] >= nums[i-1] {
			return false
		}
	}
	return true
}

func validDiffs(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	for i := 1; i < len(nums); i++ {
		diff := absInt(nums[i], nums[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
