package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	ID         int
	Population Bag
	Samples    []Bag
}

func NewGame(line string) *Game {
	idAndSamples := strings.Split(line, ": ")
	idString := strings.Split(idAndSamples[0], " ")
	id, err := strconv.Atoi(idString[1])
	if err != nil {
		panic(err)
	}

	sampleStrings := strings.Split(idAndSamples[1], ";")
	samples := make([]Bag, len(sampleStrings))
	for i, sampleString := range sampleStrings {
		sample := Bag{}

		groups := strings.Split(sampleString, ", ")
		for _, g := range groups {
			numberAndColour := strings.Fields(g)
			number, err := strconv.Atoi(numberAndColour[0])
			if err != nil {
				panic(err)
			}

			switch numberAndColour[1] {
			case "blue":
				sample.Blue = number
			case "red":
				sample.Red = number
			case "green":
				sample.Green = number
			default:
				panic(numberAndColour)
			}
		}

		samples[i] = sample
	}

	return &Game{
		ID:      id,
		Samples: samples,
	}
}

type Bag struct {
	Green int
	Red   int
	Blue  int
}

func (g *Game) InferPopulation() {
	var red int
	var green int
	var blue int

	for _, s := range g.Samples {
		if s.Red > red {
			red = s.Red
		}

		if s.Blue > blue {
			blue = s.Blue
		}

		if s.Green > green {
			green = s.Green
		}
	}

	g.Population = Bag{Red: red, Blue: blue, Green: green}
}

func (g *Game) Power() int {
	pop := g.Population
	return pop.Red * pop.Green * pop.Blue
}

func (g *Game) Possible() bool {
	pop := g.Population
	for _, s := range g.Samples {
		if s.Green > pop.Green || s.Red > pop.Red || s.Blue > pop.Blue {
			return false
		}
	}

	return true
}

func Solution(input []string) {
	partOneSolution := partOne(input)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(input)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partTwo(lines []string) int {
	var answer int

	for _, line := range lines {
		game := NewGame(line)
		game.InferPopulation()

		answer += game.Power()
	}

	return answer
}

func partOne(lines []string) int {
	var answer int

	for _, line := range lines {
		game := NewGame(line)
		game.Population = Bag{
			Red:   12,
			Green: 13,
			Blue:  14,
		}

		if game.Possible() {
			answer += game.ID
		}
	}

	return answer
}
