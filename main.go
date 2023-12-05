package main

import (
	"aoc-2023/common"
	"aoc-2023/day1"
	"aoc-2023/day2"
	"aoc-2023/day3"
	"aoc-2023/day4"
	"aoc-2023/day5"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	day, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	inputPath := fmt.Sprintf("day%d/input", day)

	switch day {
	case 1:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day1.Solution(input)
	case 2:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day2.Solution(input)
	case 3:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day3.Solution(input)
	case 4:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day4.Solution(input)
	case 5:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day5.Solution(input)
	}
}
