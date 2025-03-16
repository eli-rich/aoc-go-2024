package main

import (
	"flag"
	"fmt"

	"github.com/eli-rich/aoc-go-2024/internal/fetch"
	"github.com/eli-rich/aoc-go-2024/internal/submit"
)

func main() {
	fetchCmd := flag.Bool("fetch", false, "Fetch input for the day")
	submitCmd := flag.Bool("submit", false, "Submit answer for the day")
	part := flag.Int("part", 0, "Which part to run (1 or 2, 0 for both)")
	day := flag.Int("day", 0, "Which day to run (1-25, 0 for all available)")
	answer := flag.String("answer", "", "Answer to submit (used with -submit)")
	flag.Parse()

	// If no flags, print help:
	if !*fetchCmd && !*submitCmd {
		flag.Usage()
		return
	}

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
		if err := fetch.GetInput(*day); err != nil {
			fmt.Printf("Error fetching input for day %d: %v\n", *day, err)
			return
		}
		fmt.Printf("Fetched input for day %d\n", *day)
		return
	}
	if *submitCmd {
		if err := submit.Answer(*day, *part, *answer); err != nil {
			fmt.Printf("Error submitting answer for day %d part %d: %v\n", *day, *part, err)
			return
		}
		fmt.Printf("Submitted answer for day %d part %d\n", *day, *part)
		return
	}
}
