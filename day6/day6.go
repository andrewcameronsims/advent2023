package day6

import (
	"fmt"
	"strconv"
	"strings"
)

type Race struct {
	Duration int
	Record   int
}

func (r *Race) RecordBeaters() []int {
	var rbs []int

	for i := 1; i < r.Duration; i++ {
		dist := r.distance(i)

		if dist > r.Record {
			rbs = append(rbs, i)
		}
	}

	return rbs
}

func (r *Race) distance(timeHeld int) int {
	if timeHeld > r.Duration {
		return 0
	}

	timeRunning := r.Duration - timeHeld
	return timeHeld * timeRunning
}

func Solution(lines []string) {
	partOneSolution := partOne(lines)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(lines)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partOne(lines []string) int {
	ans := 1

	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]

	for i := 0; i < len(times); i++ {
		dur, err := strconv.Atoi(times[i])
		if err != nil {
			panic(err)
		}

		rec, err := strconv.Atoi(distances[i])
		if err != nil {
			panic(err)
		}

		race := &Race{Duration: dur, Record: rec}
		recordBeaters := race.RecordBeaters()

		ans *= len(recordBeaters)
	}

	return ans
}

func partTwo(lines []string) int {
	race := &Race{Duration: 51699878, Record: 377117112241505}
	recordBeaters := race.RecordBeaters()

	return len(recordBeaters)
}
