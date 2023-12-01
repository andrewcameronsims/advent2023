package day1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Solution(input []string) {
	partOneSolution := partOne(input)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(input)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partOne(lines []string) int {
	var answer int

	re := regexp.MustCompile(`\d`)
	for _, line := range lines {
		digits := re.FindAllString(line, -1)

		firstMatch := digits[0]
		lastMatch := digits[len(digits)-1]

		first, err := strconv.Atoi(firstMatch)
		if err != nil {
			panic(err)
		}

		last, err := strconv.Atoi(lastMatch)
		if err != nil {
			panic(err)
		}

		answer += (first*10 + last)
	}

	return answer
}

var Digits []string = []string{
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
}

var NumberWords []string = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func partTwo(lines []string) int {
	var answer int

	for _, line := range lines {
		sparseMatches := make([]int, 100)
		for i, n := range Digits {
			j := 0
			for {
				idx := strings.Index(line[j:], n)
				if idx == -1 {
					break
				}

				match := i + 1
				sparseMatches[idx+j] = match

				j += 1 + idx
			}
		}

		for i, n := range NumberWords {
			j := 0
			for {
				idx := strings.Index(line[j:], n)
				if idx == -1 {
					break
				}

				match := i + 1
				sparseMatches[idx+j] = match

				j += 1 + idx
			}
		}

		var matches []int
		for _, v := range sparseMatches {
			if v > 0 {
				matches = append(matches, v)
			}
		}

		first := matches[0]
		last := matches[len(matches)-1]

		answer += (first*10 + last)
	}

	return answer
}
