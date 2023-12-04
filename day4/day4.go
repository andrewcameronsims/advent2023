package day4

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	Id      int
	Winning []int
	Actual  []int
}

func (c *Card) Score() (int, []int) {
	var score int
	var won int

	for _, winning := range c.Winning {
		for _, actual := range c.Actual {
			if winning == actual {
				if score == 0 {
					score = 1
					won++
				} else {
					score *= 2
					won++
				}
			}
		}
	}

	cards := make([]int, won)
	for i := 1; i <= won; i++ {
		cards[i-1] = c.Id + i
	}

	return score, cards
}

func NewCard(line string) *Card {
	idAndNumbers := strings.Split(line, ": ")

	id := strings.Split(idAndNumbers[0], "Card")[1]
	winningAndActual := strings.Split(idAndNumbers[1], " | ")

	winning := strings.Fields(winningAndActual[0])
	actual := strings.Fields(winningAndActual[1])

	idInt, err := strconv.Atoi(strings.TrimSpace(id))
	if err != nil {
		panic(err)
	}

	card := &Card{
		Id:      idInt,
		Winning: make([]int, len(winning)),
		Actual:  make([]int, len(actual)),
	}

	for i, w := range winning {
		win, err := strconv.Atoi(w)
		if err != nil {
			panic(err)
		}

		card.Winning[i] = win
	}

	for i, a := range actual {
		act, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}

		card.Actual[i] = act
	}

	return card
}

func Solution(lines []string) {
	partOneSolution := partOne(lines)
	fmt.Printf("Part One: %d\n", partOneSolution)

	partTwoSolution := partTwo(lines)
	fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partOne(lines []string) int {
	var result int

	for _, line := range lines {
		card := NewCard(line)
		score, _ := card.Score()
		result += score
	}

	return result
}

func partTwo(lines []string) int {
	var result int
	cs := make(map[int]int)

	for _, line := range lines {
		card := NewCard(line)

		times, ok := cs[card.Id]
		if !ok {
			cs[card.Id] = 1
			times = 1
		}
		for i := 0; i < times; i++ {
			_, copies := card.Score()

			for _, copy := range copies {
				curr, ok := cs[copy]
				if !ok {
					cs[copy] = 2
				} else {
					cs[copy] = curr + 1
				}
			}
		}
	}

	for _, v := range cs {
		result += v
	}

	return result
}
