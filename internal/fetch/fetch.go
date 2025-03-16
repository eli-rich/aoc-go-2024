package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/eli-rich/aoc-go-2024/internal/utils"
)

func GetInput(day int) error {
	dayDir := filepath.Join("internal", "days", "day"+strconv.Itoa(day))
	if err := os.Mkdir(dayDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory for day %d: %w", day, err)
	}
	outputPath := filepath.Join(dayDir, "input.txt")
	exists := utils.CheckPathExists(outputPath)
	if exists {
		return fmt.Errorf("input file for day %d already exists", day)
	}

	// fetch the input from the Advent of Code website
	sessionCookie := os.Getenv("AOC_COOKIE")
	if sessionCookie == "" {
		return fmt.Errorf("AOC_COOKIE environment variable is not set")
	}

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch input: %s", resp.Status)
	}
	// write the response body to the output file
	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer out.Close()

	// copy the response body to the output file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to output file: %w", err)
	}
	return nil
}
