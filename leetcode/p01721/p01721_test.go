package p01721

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	head []*ListNode
	k    int
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
				head: []*ListNode{
					{Val: 1},
					{Val: 2},
					{Val: 3},
					{Val: 4},
					{Val: 5},
				},
				k: 2,
			},
			output{
				ans: []int{1, 4, 3, 2, 5},
			},
		},
		{
			input{
				head: []*ListNode{
					{Val: 7},
					{Val: 9},
					{Val: 6},
					{Val: 6},
					{Val: 7},
					{Val: 8},
					{Val: 3},
					{Val: 0},
					{Val: 9},
					{Val: 5},
				},
				k: 5,
			},
			output{
				ans: []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5},
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.head)-1; i++ {
			data.i.head[i].Next = data.i.head[i+1]
		}

		head := swapNodes(data.i.head[0], data.i.k)
		res := []int{}
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		ast.Equal(data.o.ans, res)
	}
}
