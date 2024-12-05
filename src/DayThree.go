package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func DayThree() {
	f, err := os.Open("inputs/DayThreeInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var matches []string
	answer := 0

	// Setup scanner and regex
	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	// Collect matches
	for scanner.Scan() {
		line := scanner.Text()
		lineMatches := re.FindAllString(line, -1)
		matches = append(matches, lineMatches...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	re = regexp.MustCompile(`\d+,\d+`)
	for _, mulString := range matches {
		numString := re.FindString(mulString)
		nums := strings.Split(numString, ",")
		numOne, _ := strconv.Atoi(nums[0])
		numTwo, _ := strconv.Atoi(nums[1])
		answer += numOne * numTwo
	}

	fmt.Println(answer)
}

func DayThreePtTwo() {
	// Read the content of the file
	content, err := os.ReadFile("inputs/DayThreeInput.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// Convert the content to a string
	text := string(content)
	re := regexp.MustCompile(`do\(\)`)
	// Split the content at 'do()'
	lines := re.Split(text, -1)

	var correctLines []string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Check if the line contains 'don't()' and cut off after it
		if idx := strings.Index(line, "don't()"); idx != -1 {
			// Keep everything before 'don't()'
			line = line[:idx]
		}

		// Print the cleaned line, only if it isn't empty
		if line != "" {
			correctLines = append(correctLines, line)
		}
	}

	var matches []string
	answer := 0

	re = regexp.MustCompile(`mul\(\d+,\d+\)`)

	// Collect matches
	for _, line := range correctLines {
		lineMatches := re.FindAllString(line, -1)
		matches = append(matches, lineMatches...)
	}

	re = regexp.MustCompile(`\d+,\d+`)
	for _, mulString := range matches {
		numString := re.FindString(mulString)
		nums := strings.Split(numString, ",")
		numOne, _ := strconv.Atoi(nums[0])
		numTwo, _ := strconv.Atoi(nums[1])
		answer += numOne * numTwo
	}

	fmt.Println(answer)
}
