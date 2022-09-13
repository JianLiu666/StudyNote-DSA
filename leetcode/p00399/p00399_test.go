package p00399

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	equations [][]string
	values    []float64
	queries   [][]string
}

type output struct {
	ans []float64
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{
				equations: [][]string{{"x1", "x2"}, {"x2", "x3"}, {"x1", "x4"}, {"x2", "x5"}},
				values:    []float64{3.0, 0.5, 3.4, 5.6},
				queries:   [][]string{{"x2", "x4"}, {"x1", "x5"}, {"x1", "x3"}, {"x5", "x5"}, {"x5", "x1"}, {"x3", "x4"}, {"x4", "x3"}, {"x6", "x6"}, {"x0", "x0"}},
			},
			output{
				ans: []float64{1.13333, 16.80000, 1.50000, 1.00000, 0.05952, 2.26667, 0.44118, -1.00000, -1.00000},
			},
		},
		{
			input{
				equations: [][]string{{"a", "b"}, {"e", "f"}, {"b", "e"}},
				values:    []float64{3.4, 1.4, 2.3},
				queries:   [][]string{{"b", "a"}, {"a", "f"}, {"f", "f"}, {"e", "e"}, {"c", "c"}, {"a", "c"}, {"f", "e"}},
			},
			output{
				ans: []float64{0.29412, 10.94800, 1.00000, 1.00000, -1.00000, -1.00000, 0.71429},
			},
		},
		{
			input{
				equations: [][]string{{"a", "b"}, {"b", "c"}},
				values:    []float64{2.0, 3.0},
				queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			},
			output{
				ans: []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
			},
		},
		{
			input{
				equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
				values:    []float64{1.5, 2.5, 5.0},
				queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			},
			output{
				ans: []float64{3.75000, 0.40000, 5.00000, 0.20000},
			},
		},
		{
			input{
				equations: [][]string{{"a", "b"}},
				values:    []float64{0.5},
				queries:   [][]string{{"a", "b"}, {"b", "a"}},
			},
			output{
				ans: []float64{0.50000, 2.00000},
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, calcEquation(data.i.equations, data.i.values, data.i.queries), idx+1)
	}
}
