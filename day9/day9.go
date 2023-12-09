package day9

import (
	"aoc-2023/common"
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
	var ans int

	for _, line := range lines {
		hist := NewHistory(line)
		hist.Derive()

		ans += hist.ExtrapolateForwards()
	}

	return ans
}

func partTwo(lines []string) int {
	var ans int

	for _, line := range lines {
		hist := NewHistory(line)
		hist.Derive()

		ans += hist.ExtrapolateBack()
	}

	return ans
}

type History struct {
	Original []int
	Derived  [][]int
}

func NewHistory(line string) *History {
	nums := strings.Fields(line)
	original := make([]int, len(nums))
	for i, n := range nums {
		original[i] = common.Int(n)
	}

	return &History{
		Original: original,
	}
}

func (h *History) ExtrapolateForwards() int {
	for i := len(h.Derived) - 1; i >= 0; i-- {
		if i == len(h.Derived)-1 {
			h.Derived[i] = append(h.Derived[i], 0)
			continue
		}

		nextValue := h.Derived[i][len(h.Derived[i])-1] + h.Derived[i+1][len(h.Derived[i+1])-1]
		h.Derived[i] = append(h.Derived[i], nextValue)
	}

	return h.Original[len(h.Original)-1] + h.Derived[0][len(h.Derived[0])-1]
}

func (h *History) ExtrapolateBack() int {
	for i := len(h.Derived) - 1; i >= 0; i-- {
		if i == len(h.Derived)-1 {
			h.Derived[i] = append([]int{0}, h.Derived[i]...)
			continue
		}

		nextValue := h.Derived[i][0] - h.Derived[i+1][0]
		h.Derived[i] = append([]int{nextValue}, h.Derived[i]...)
	}

	return h.Original[0] - h.Derived[0][0]
}

func (h *History) Derive() {
	current := h.Original

	for !allZeros(current) {
		var derived []int

		for i := 0; i < len(current); i++ {
			if i+1 >= len(current) {
				break
			}

			change := current[i+1] - current[i]
			derived = append(derived, change)
		}

		h.Derived = append(h.Derived, derived)
		current = h.Derived[len(h.Derived)-1]
	}
}

func allZeros(ns []int) bool {
	for _, n := range ns {
		if n != 0 {
			return false
		}
	}

	return true
}
