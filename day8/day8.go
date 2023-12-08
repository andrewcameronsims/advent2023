package day8

import (
	"fmt"
	"strings"
)

func Solution(lines []string) {
	partOneSolution := partOne(lines)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(lines)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partOne(lines []string) int {
	turns := lines[0]
	paths := makeGraph(lines[2:])

	var i int
	curr := "AAA"
	for {
		idx := i % len(turns)
		t := turns[idx]
		if t == 'L' {
			curr = paths[curr][0]
		} else if t == 'R' {
			curr = paths[curr][1]
		} else {
			panic("oops")
		}

		i++

		if curr == "ZZZ" {
			break
		}
	}

	return i
}

func partTwo(lines []string) int {
	turns := lines[0]
	paths := makeGraph(lines[2:])

	var currs []string
	for key := range paths {
		if key[2] == 'A' {
			currs = append(currs, key)
		}
	}

	// get the independent cycle times for all the starting nodes
	var cycleTimes []int
	for _, curr := range currs {
		var i int

		for {
			idx := i % len(turns)
			t := turns[idx]
			if t == 'L' {
				curr = paths[curr][0]
			} else if t == 'R' {
				curr = paths[curr][1]
			} else {
				panic("oops")
			}

			i++

			if curr[2] == 'Z' {
				break
			}
		}

		cycleTimes = append(cycleTimes, i)
	}

	// when do the cycles line up?
	return leastCommonMultiple(cycleTimes[0], cycleTimes[1], cycleTimes[2:]...)
}

func makeGraph(lines []string) map[string][]string {
	g := make(map[string][]string)

	for _, line := range lines {
		srcAndDsts := strings.Split(line, " = ")
		dsts := strings.Split(strings.Trim(srcAndDsts[1], "()"), ", ")

		headVal := srcAndDsts[0]
		leftVal := dsts[0]
		rightVal := dsts[1]

		g[headVal] = []string{leftVal, rightVal}
	}

	return g
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func leastCommonMultiple(a, b int, ns ...int) int {
	result := a * b / greatestCommonDivisor(a, b)

	for i := 0; i < len(ns); i++ {
		result = leastCommonMultiple(result, ns[i])
	}

	return result
}
