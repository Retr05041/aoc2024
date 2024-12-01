package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func processDayOneInput() ([]int, []int) {
	// open file
	f, err := os.Open("inputs/DayOneInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	var left = make([]int, 1000)
	var right = make([]int, 1000)

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	i := 0
	for scanner.Scan() {
		// do something with a line
		line := strings.Split(scanner.Text(), "   ")
		left[i], err = strconv.Atoi(strings.ReplaceAll(line[0], " ", ""))
		if err != nil {
			fmt.Println(err)
		}

		right[i], err = strconv.Atoi(strings.ReplaceAll(line[1], " ", ""))
		if err != nil {
			fmt.Println(err)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return left, right
}

func DayOne() {
	left, right := processDayOneInput()
	sort.Ints(left)
	sort.Ints(right)

	answer := 0

	for i := range left {
		if left[i] > right[i] {
			answer += left[i] - right[i]
		} else if left[i] < right[i] {
			answer += right[i] - left[i]
		} else {
			answer += 0
		}
	}

	fmt.Println(answer)
}

func DayOnePtTwo() {
	left, right := processDayOneInput()

	numOfXInY := 0
	answer := 0
	for i := range left {
		for j := range right {
			if left[i] == right[j] {
				numOfXInY++
			}
		}

		answer += left[i] * numOfXInY
		numOfXInY = 0
	}

	fmt.Println(answer)

}
