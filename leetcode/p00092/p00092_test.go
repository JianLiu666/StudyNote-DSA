package p00092

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	head  []*ListNode
	left  int
	right int
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
				head:  []*ListNode{{Val: 1}, {Val: 2}, {Val: 3}, {Val: 4}, {Val: 5}},
				left:  2,
				right: 4,
			},
			output{
				ans: []int{1, 4, 3, 2, 5},
			},
		},
		{
			input{
				head:  []*ListNode{{Val: 5}},
				left:  1,
				right: 1,
			},
			output{
				ans: []int{5},
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.head)-1; i++ {
			data.i.head[i].Next = data.i.head[i+1]
		}

		head := reverseBetween(data.i.head[0], data.i.left, data.i.right)
		res := []int{}
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		ast.Equal(data.o.ans, res)
	}
}
