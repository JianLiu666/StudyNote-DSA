package p00412

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
}

type output struct {
	ans []string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{n: 3},
			output{ans: []string{"1", "2", "Fizz"}},
		},
		{
			input{n: 5},
			output{ans: []string{"1", "2", "Fizz", "4", "Buzz"}},
		},
		{
			input{n: 15},
			output{ans: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, fizzBuzz(data.i.n))
	}
}
