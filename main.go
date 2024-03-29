package main

import (
	"aoc-2023/common"
	"aoc-2023/day1"
	"aoc-2023/day10"
	"aoc-2023/day11"
	"aoc-2023/day2"
	"aoc-2023/day3"
	"aoc-2023/day4"
	"aoc-2023/day5"
	"aoc-2023/day6"
	"aoc-2023/day7"
	"aoc-2023/day8"
	"aoc-2023/day9"
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
	case 6:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day6.Solution(input)
	case 7:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day7.Solution(input)
	case 8:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day8.Solution(input)
	case 9:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day9.Solution(input)
	case 10:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day10.Solution(input)
	case 11:
		input, err := common.ReadLinesFromInput(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		day11.Solution(input)
	}
}
