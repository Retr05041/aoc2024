package src

import (
	"bufio"
	"fmt"
	"lib"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func processDayTwoInput() [][]int {
	// open file
	f, err := os.Open("inputs/DayTwoInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	var input = make([][]int, 1000)

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	i := 0
	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), " ")
		input[i] = make([]int, len(line))

		for j, s := range line {
			input[i][j], _ = strconv.Atoi(s)
		}

		i++
	}

	return input
}

func dayTwo_isSafe(report []int) bool {

	ascending := true
	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			diff := report[i+1] - report[i]
			if diff < 1 || diff > 3 {
				ascending = false
			}
		} else {
			ascending = false
		}
	}

	descending := true
	for i := 0; i < len(report)-1; i++ {
		if report[i] > report[i+1] {
			diff := report[i] - report[i+1]
			if diff < 1 || diff > 3 {
				descending = false
			}
		} else {
			descending = false
		}
	}

	if ascending || descending {
		return true
	}
	return false
}

func DayTwo() {
	reports := processDayTwoInput()

	answer := 0
	for i := range reports {
		if dayTwo_isSafe(reports[i]) {
			answer++
		}
	}

	fmt.Println(answer)
}

func dayTwoPtTwo_ascending(report []int, tolerance int) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			diff := report[i+1] - report[i]
			if diff < 1 || diff > 3 {
				if tolerance < 1 {
					tolerance++
					tmpReport := make([]int, len(report))
					copy(tmpReport, report)
					tmpReport = lib.Remove(tmpReport, i+1)
					if dayTwoPtTwo_ascending(tmpReport, 1) {
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			}
			// fmt.Println("Diff: " + strconv.Itoa(diff) + " - (" + strconv.Itoa(report[i]) + "," + strconv.Itoa(report[i+1]) + ")")
		} else {
			if tolerance < 1 {
				tolerance++
				tmpReport := make([]int, len(report))
				copy(tmpReport, report)
				tmpReport = lib.Remove(tmpReport, i+1)
				if dayTwoPtTwo_ascending(tmpReport, 1) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	return true
}

func dayTwoPtTwo_descending(report []int, tolerance int) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i] > report[i+1] {
			diff := report[i] - report[i+1]
			if diff < 1 || diff > 3 {
				if tolerance < 1 {
					tolerance++
					tmpReport := make([]int, len(report))
					copy(tmpReport, report)
					tmpReport = lib.Remove(tmpReport, i+1)
					if dayTwoPtTwo_descending(tmpReport, 1) {
						return true
					} else {
						return false
					}
				} else {
					if tolerance < 1 {
					}
					return false
				}
			}
		} else {
			if tolerance < 1 {
				tolerance++
				tmpReport := make([]int, len(report))
				copy(tmpReport, report)
				tmpReport = lib.Remove(tmpReport, i+1)
				if dayTwoPtTwo_descending(tmpReport, 1) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	return true
}

func dayTwoPtTwo_isSafe(report []int) bool {
	copy1 := make([]int, len(report))
	copy(copy1, report)
	copy2 := make([]int, len(report))
	copy(copy2, report)
	ascending := dayTwoPtTwo_ascending(copy1, 0)
	if !ascending {
		slices.Reverse(copy1)
		ascending = dayTwoPtTwo_ascending(copy1, 0)
	}
	descending := dayTwoPtTwo_descending(copy2, 0)
	if !descending {
		slices.Reverse(copy2)
		descending = dayTwoPtTwo_descending(copy2, 0)
	}

	if ascending || descending {
		return true
	}
	return false
}

func DayTwoPtTwo() {
	reports := processDayTwoInput()

	answer := 0
	for i := range reports {
		if dayTwoPtTwo_isSafe(reports[i]) {
			answer++
		}
	}

	fmt.Println(answer)
}
