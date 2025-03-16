package submit

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Result patterns to determine if the submission was successful or not
var (
	correctPattern = regexp.MustCompile(`That's the right answer`)
	tooLowPattern  = regexp.MustCompile(`Your answer is too low`)
	tooHighPattern = regexp.MustCompile(`Your answer is too high`)
	wrongPattern   = regexp.MustCompile(`That's not the right answer`)
	waitPattern    = regexp.MustCompile(`You gave an answer too recently`)
)

func Answer(day, part int, answer string) error {
	sessionCookie := os.Getenv("AOC_COOKIE")
	if sessionCookie == "" {
		return fmt.Errorf("AOC_COOKIE environment variable is not set")
	}

	// Build form data
	formData := url.Values{
		"level":  {strconv.Itoa(part)},
		"answer": {answer},
	}

	// Create the request
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/answer", day)
	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
		return fmt.Errorf("failed to submit answer: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	responseText := string(body)
	switch {
	case correctPattern.MatchString(responseText):
		return nil // Correct answer
	case tooLowPattern.MatchString(responseText):
		return fmt.Errorf("your answer is too low")
	case tooHighPattern.MatchString(responseText):
		return fmt.Errorf("your answer is too high")
	case wrongPattern.MatchString(responseText):
		return fmt.Errorf("that's not the right answer")
	case waitPattern.MatchString(responseText):
		// find wait time
		waitTimePattern := regexp.MustCompile(`You have (\d+m \d+s) left to wait`)
		matches := waitTimePattern.FindStringSubmatch(responseText)
		if len(matches) > 1 {
			return fmt.Errorf("you need to wait %s before submitting again", matches[1])
		}
		return fmt.Errorf("you are submitting too frequently")
	default:
		return fmt.Errorf("unexpected response, check advent of code website")
	}
}
