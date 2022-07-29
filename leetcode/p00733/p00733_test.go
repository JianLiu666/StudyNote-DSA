package p00733

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	image [][]int
	sr    int
	sc    int
	color int
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
			input{image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
				sr: 1, sc: 1, color: 2},
			output{ans: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			}},
		},
		{
			input{image: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
				sr: 0, sc: 0, color: 0},
			output{ans: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			}},
		},
	}

	for idx, data := range tds {
		copyright := make([][]int, len(data.i.image))
		for i := 0; i < len(data.i.image); i++ {
			arr := make([]int, len(data.i.image[i]))
			copy(arr, data.i.image[i])
			copyright[i] = arr
		}
		ast.Equal(data.o.ans, floodFill_dfs(copyright, data.i.sr, data.i.sc, data.i.color), fmt.Sprintf("dfs case %d", idx+1))
	}
}
