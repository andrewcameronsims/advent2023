package day3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3_PartOne(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		sum   int
	}{
		{
			desc: "",
			input: `
				467..114..
				...*......
				..35..633.
				......#...
				617*......
				.....+.58.
				..592.....
				......755.
				...$.*....
				.664.598..
			`,
			sum: 4361,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lines := strings.Fields(tC.input)
			matrix := make([][]rune, len(lines))
			for i, line := range lines {
				matrix[i] = []rune(line)
			}

			sch := &Schematic{
				Matrix: matrix,
			}
			sch.FindParts()

			sum := sch.SumParts()

			assert.Equal(t, tC.sum, sum)
		})
	}
}

func TestDay3_PartTwo(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		sum   int
	}{
		{
			desc: "",
			input: `
				467..114..
				...*......
				..35..633.
				......#...
				617*......
				.....+.58.
				..592.....
				......755.
				...$.*....
				.664.598..
			`,
			sum: 467835,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lines := strings.Fields(tC.input)
			matrix := make([][]rune, len(lines))
			for i, line := range lines {
				matrix[i] = []rune(line)
			}

			sch := &Schematic{
				Matrix: matrix,
			}
			sch.FindParts()

			sum := sch.GearRatios()

			assert.Equal(t, tC.sum, sum)
		})
	}
}
