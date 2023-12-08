package day8

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8_PartOne(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		ans   int
	}{
		{
			desc: "",
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			ans: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lines := strings.Split(tC.input, "\n")
			ans := partOne(lines)

			assert.Equal(t, tC.ans, ans)
		})
	}
}
