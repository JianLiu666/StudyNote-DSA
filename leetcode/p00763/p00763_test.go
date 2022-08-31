package p00763

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
}

type output struct {
	ans []int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{s: "ababcbacadefegdehijhklij"},
			output{ans: []int{9, 7, 8}},
		},
		{
			input{s: "eccbbbbdec"},
			output{ans: []int{10}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, partitionLabels(data.i.s), idx+1)
	}
}
