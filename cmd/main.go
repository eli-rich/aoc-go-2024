package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/eli-rich/aoc-go-2024/internal/days/day1"
	"github.com/eli-rich/aoc-go-2024/internal/fetch"
	"github.com/eli-rich/aoc-go-2024/internal/submit"
)

type daySolver interface {
	Part1() (string, error)
	Part2() (string, error)
}

var solvers = map[int]daySolver{
	1: day1.Solver{},
}

func runDay(day, part int) {
	solver, exists := solvers[day]
	if !exists {
		fmt.Printf("Solution for day %d not implemented yet\n", day)
		return
	}
	if part == 0 || part == 1 {
		start := time.Now()
		answer, err := solver.Part1()
		elapsed := time.Since(start)
		if err != nil {
			fmt.Printf("Part 1 error: %v\n", err)
		} else {
			fmt.Printf("Part 1: %s (took %s)\n", answer, elapsed)
		}
	}
	if part == 0 || part == 2 {
		start := time.Now()
		answer, err := solver.Part2()
		elapsed := time.Since(start)
		if err != nil {
			fmt.Printf("Part 2 error: %v\n", err)
		} else {
			fmt.Printf("Part 2: %s (took %s)\n", answer, elapsed)
		}
	}
}

func submitAnswer(day, part int, answer string) {
	fmt.Printf("Submitting answer for day %d part %d: %s\n", day, part, answer)
	result := submit.Answer(day, part, answer)
	fmt.Println(result)
}

func main() {
	fetchCmd := flag.Bool("fetch", false, "Fetch input for the day")
	submitCmd := flag.Bool("submit", false, "Submit answer for the day")
	part := flag.Int("part", 0, "Which part to run (1 or 2, 0 for both)")
	day := flag.Int("day", 0, "Which day to run (1-25, 0 for all available)")
	answer := flag.String("answer", "", "Answer to submit (used with -submit)")
	flag.Parse()

	// Validate command line arguments
	if *fetchCmd && *submitCmd {
		fmt.Println("Error: Cannot fetch and submit at the same time")
		return
	}
	// Validate day argument
	if *day < 0 || *day > 25 {
		fmt.Println("Error: day must be between 0 and 25")
		return
	}
	// Validate part argument
	if *part < 0 || *part > 2 {
		fmt.Println("Error: part must be between 0 and 2")
		return
	}
	// Validate answer argument
	if *submitCmd && *answer == "" {
		fmt.Println("Error: answer must be specified when submitting")
		return
	}

	// Execute the appropriate command
	if *fetchCmd {
		if *day == 0 {
			fmt.Println("Error: day must be specified when fetching")
		}
		if err := fetch.GetInput(*day); err != nil {
			fmt.Printf("Error fetching input for day %d: %v\n", *day, err)
			return
		}
		fmt.Printf("Fetched input for day %d\n", *day)
		return
	}
	if *submitCmd {
		if *day == 0 || *part == 0 {
			fmt.Println("Error: day and part must be specified when submitting")
			return
		}
		submitAnswer(*day, *part, *answer)
		return
	}

	// Run days
	if *day == 0 {
		// Run all available days
		for d := range solvers {
			fmt.Printf("--- Day %d ---\n", d)
			runDay(d, *part)
			fmt.Println()
		}
	} else {
		// Run specific day
		fmt.Printf("--- Day %d ---\n", *day)
		runDay(*day, *part)
	}
}
