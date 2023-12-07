package day7

import (
	"aoc-2023/common"
	"fmt"
	"sort"
	"strings"
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var Cards = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 0,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

type Hand struct {
	Type         HandType
	Cards        map[int]int
	CardsOrdered []int
	Bid          int
}

func (h *Hand) inferType() {
	cards := h.Cards

	var pairs int
	var fives int
	var fours int
	var threes int
	var jokers int

	for card, held := range cards {
		if card == 0 {
			jokers = held
		}

		if held == 5 && card != 0 {
			fives++
		}

		if held == 4 && card != 0 {
			fours++
		}

		if held == 3 && card != 0 {
			threes++
		}

		if held == 2 && card != 0 {
			pairs++
		}
	}

	if fives == 1 || (fours == 1 && jokers == 1) || (jokers == 2 && threes == 1) || (jokers == 3 && pairs == 1) || (jokers == 4) || (jokers == 5) {
		h.Type = FiveOfAKind
		return
	}

	if fours == 1 || (threes == 1 && jokers == 1) || (jokers == 2 && pairs == 1) || (jokers == 3) {
		h.Type = FourOfAKind
		return
	}

	if threes == 1 && pairs == 1 || (pairs == 2 && jokers == 1) {
		h.Type = FullHouse
		return
	}

	if threes == 1 || (pairs == 1 && jokers == 1) || (jokers == 2) {
		h.Type = ThreeOfAKind
		return
	}

	if pairs == 2 {
		h.Type = TwoPair
		return
	}

	if pairs == 1 || jokers == 1 {
		h.Type = OnePair
		return
	}

	h.Type = HighCard
}

func NewHand(raw string) *Hand {
	cardsAndBid := strings.Split(raw, " ")
	cardsRaw := []rune(cardsAndBid[0])
	bid := common.Int(cardsAndBid[1])

	cards := make(map[int]int)
	ordered := make([]int, len(cardsRaw))
	for i := range cardsRaw {

		card, ok := Cards[cardsRaw[i]]
		if !ok {
			panic(cardsRaw[i])
		}

		cards[card] += 1
		ordered[i] = card
	}

	h := Hand{
		Cards:        cards,
		CardsOrdered: ordered,
		Bid:          bid,
	}
	h.inferType()

	return &h
}

type ByRank []*Hand

func (hs ByRank) Len() int      { return len(hs) }
func (hs ByRank) Swap(i, j int) { hs[i], hs[j] = hs[j], hs[i] }
func (hs ByRank) Less(i, j int) bool {
	left := hs[i]
	right := hs[j]

	if left.Type != right.Type {
		return left.Type < right.Type
	}

	for c := 0; c < 5; c++ {
		if left.CardsOrdered[c] != right.CardsOrdered[c] {
			return left.CardsOrdered[c] < right.CardsOrdered[c]
		}
	}

	panic("can't order")
}

func Solution(lines []string) {
	partOneSolution := partOne(lines)
	fmt.Printf("Part One: %d\n", partOneSolution)

	// partTwoSolution := partTwo(lines)
	// fmt.Printf("Part Two: %d\n", partTwoSolution)
}

func partOne(lines []string) int {
	var winnings int

	hands := make(ByRank, len(lines))
	for i, line := range lines {
		hands[i] = NewHand(line)
	}

	sort.Sort(hands)

	for i, h := range hands {
		fmt.Printf("%+v\n", *h)

		winnings += (h.Bid * (i + 1))
	}

	return winnings
}
