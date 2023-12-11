package day11

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay11_PartOne(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expanded string
	}{
		{
			desc: "",
			input: `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`,
			expanded: `....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lines := strings.Split(tC.input, "\n")
			gal := NewUniverse(lines)

			gal.Expand()

			expanded := strings.TrimSpace(gal.String())

			assert.Equal(t, tC.expanded, expanded)
		})
	}
}

func TestDay11_PartTwo(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expanded string
	}{
		{
			desc: "",
			input: `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lines := strings.Split(tC.input, "\n")
			gal := NewUniverse(lines)
			gal.MapExpansionBoundaries()

		})
	}
}
