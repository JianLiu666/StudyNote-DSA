package p00234

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	head []*ListNode
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
			input{
				head: []*ListNode{{Val: 1}, {Val: 2}, {Val: 2}, {Val: 1}},
			},
			output{
				ans: true,
			},
		},
		{
			input{
				head: []*ListNode{{Val: 1}, {Val: 2}},
			},
			output{
				ans: false,
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.head)-1; i++ {
			data.i.head[i].Next = data.i.head[i+1]
		}
		ast.Equal(data.o.ans, isPalindrome(data.i.head[0]))
	}
}
