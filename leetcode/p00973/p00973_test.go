package p00973

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	points [][]int
	k      int
}

type output struct {
	ans [][]int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{points: [][]int{{1, 3}, {-2, 2}}, k: 1},
			output{ans: [][]int{{-2, 2}}},
		},
		{
			input{points: [][]int{{3, 3}, {5, -1}, {-2, 4}}, k: 2},
			output{ans: [][]int{{3, 3}, {-2, 4}}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, kClosest(data.i.points, data.i.k), idx+1)
	}
}
