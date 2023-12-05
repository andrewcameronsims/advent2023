package day5

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string
		lowest int
	}{
		{
			desc: "",
			input: `seeds: 79 14 55 13

			seed-to-soil map:
			50 98 2
			52 50 48

			soil-to-fertilizer map:
			0 15 37
			37 52 2
			39 0 15

			fertilizer-to-water map:
			49 53 8
			0 11 42
			42 0 7
			57 7 4

			water-to-light map:
			88 18 7
			18 25 70

			light-to-temperature map:
			45 77 23
			81 45 19
			68 64 13

			temperature-to-humidity map:
			0 69 1
			1 0 69

			humidity-to-location map:
			60 56 37
			56 93 4
			`,
			lowest: 35,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lines := strings.Split(tC.input, "\n")

			seedWords := strings.Split(lines[0], " ")[1:]
			seeds := make([]int, len(seedWords))
			for i, sw := range seedWords {
				seed, err := strconv.Atoi(sw)
				if err != nil {
					panic(err)
				}

				seeds[i] = seed
			}

			al := NewAlmanac(lines)

			loc := al.SeedToLocation(13)

			fmt.Printf("al: %v\n", al)

			assert.Equal(t, tC.lowest, loc)
		})
	}
}
