package p01079

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	tiles string
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
			input{tiles: "AAB"},
			output{ans: 8},
		},
		{
			input{tiles: "AAABBC"},
			output{ans: 188},
		},
		{
			input{tiles: "V"},
			output{ans: 1},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, numTilePossibilities_backtracking(data.i.tiles), fmt.Sprintf("Backtracking: %v", idx+1))
		ast.Equal(data.o.ans, numTilePossibilities_dfs(data.i.tiles), fmt.Sprintf("Depth-First Search: %v", idx+1))
	}
}
