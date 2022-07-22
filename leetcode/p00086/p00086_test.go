package p00086

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	head []*ListNode
	x    int
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
			input{
				head: []*ListNode{{Val: 1}, {Val: 4}, {Val: 3}, {Val: 2}, {Val: 5}, {Val: 2}},
				x:    3,
			},
			output{
				ans: []int{1, 2, 2, 4, 3, 5},
			},
		},
		{
			input{
				head: []*ListNode{{Val: 2}, {Val: 1}},
				x:    2,
			},
			output{
				ans: []int{1, 2},
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.head)-1; i++ {
			data.i.head[i].Next = data.i.head[i+1]
		}

		head := partition(data.i.head[0], data.i.x)
		res := []int{}
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		ast.Equal(data.o.ans, res)
	}
}
