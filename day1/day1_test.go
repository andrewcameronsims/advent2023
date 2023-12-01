package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		desc   string
		lines  []string
		answer int
	}{
		{
			desc:   "",
			lines:  []string{"8twosvdmcntf1hfive393"},
			answer: 83,
		},
		{
			desc:   "",
			lines:  []string{"honemkmbfbnlhtbq19twonekbp"},
			answer: 11,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := partTwo(tC.lines)
			assert.Equal(t, tC.answer, got)
		})
	}
}
