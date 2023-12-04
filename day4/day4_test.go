package day4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4(t *testing.T) {
	testCases := []struct {
		desc   string
		cards  string
		copies map[int]int
	}{
		{
			desc: "",
			cards: `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
				Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
				Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
				Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
				Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
				Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
			copies: map[int]int{
				1: 1,
				2: 2,
				3: 4,
				4: 8,
				5: 14,
				6: 1,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var result int
			cs := make(map[int]int)

			lines := strings.Split(tC.cards, "\n")

			for _, line := range lines {
				card := NewCard(line)

				times, ok := cs[card.Id]
				if !ok {
					cs[card.Id] = 1
					times = 1
				}
				for i := 0; i < times; i++ {
					score, copies := card.Score()
					result += score

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

			assert.Equal(t, tC.copies, cs)
		})
	}
}
