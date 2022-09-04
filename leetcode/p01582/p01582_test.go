package p01582

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	mat [][]int
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
			input{mat: [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}},
			output{ans: 1},
		},
		{
			input{mat: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
			output{ans: 3},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, numSpecial(data.i.mat), fmt.Sprintf("Brute Force: %v", idx+1))
		ast.Equal(data.o.ans, numSpecial_min_sc(data.i.mat), fmt.Sprintf("Minimize Space Complexity: %v", idx+1))
	}
}
