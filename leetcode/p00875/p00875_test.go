package p00875

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	piles []int
	h     int
}

type output struct {
	ans int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{piles: []int{3, 6, 7, 11}, h: 8},
			output{ans: 4},
		},
		{
			input{piles: []int{30, 11, 23, 4, 20}, h: 5},
			output{ans: 30},
		},
		{
			input{piles: []int{30, 11, 23, 4, 20}, h: 6},
			output{ans: 23},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, minEatingSpeed(data.i.piles, data.i.h), idx+1)
	}
}
