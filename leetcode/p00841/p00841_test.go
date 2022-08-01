package p00841

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	rooms [][]int
}

type output struct {
	ans bool
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{rooms: [][]int{{1}, {2}, {3}, {}}},
			output{ans: true},
		},
		{
			input{rooms: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}},
			output{ans: false},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, canVisitAllRooms(data.i.rooms), idx+1)
	}
}
