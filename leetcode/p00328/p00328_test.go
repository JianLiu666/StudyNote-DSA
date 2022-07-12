package p00328

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	head []*ListNode
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
				head: []*ListNode{{Val: 1}, {Val: 2}, {Val: 3}, {Val: 4}, {Val: 5}},
			},
			output{
				ans: []int{1, 3, 5, 2, 4},
			},
		},
		{
			input{
				head: []*ListNode{{Val: 2}, {Val: 1}, {Val: 3}, {Val: 5}, {Val: 6}, {Val: 4}, {Val: 7}},
			},
			output{
				ans: []int{2, 3, 6, 7, 1, 5, 4},
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.head)-1; i++ {
			data.i.head[i].Next = data.i.head[i+1]
		}

		head := oddEvenList(data.i.head[0])
		res := []int{}
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		ast.Equal(data.o.ans, res)
	}
}
