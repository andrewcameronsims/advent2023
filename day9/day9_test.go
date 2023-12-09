package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay9_PartOne(t *testing.T) {
	testCases := []struct {
		desc    string
		history *History
		derived [][]int
		extra   int
	}{
		{
			desc: "",
			history: &History{
				Original: []int{0, 3, 6, 9, 12, 15},
			},
			derived: [][]int{
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
			},
			extra: 18,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.history.Derive()
			assert.Equal(t, tC.derived, tC.history.Derived)

			extra := tC.history.ExtrapolateForwards()
			assert.Equal(t, tC.extra, extra)
		})
	}
}

func TestDay9_PartTwo(t *testing.T) {
	testCases := []struct {
		desc    string
		history *History
		derived [][]int
		extra   int
	}{
		{
			desc: "",
			history: &History{
				Original: []int{0, 3, 6, 9, 12, 15},
			},
			derived: [][]int{
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
			},
			extra: -3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.history.Derive()
			assert.Equal(t, tC.derived, tC.history.Derived)

			extra := tC.history.ExtrapolateBack()
			assert.Equal(t, tC.extra, extra)
		})
	}
}
