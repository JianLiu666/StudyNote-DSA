package p00237

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	head []*ListNode
	idx  int
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
				head: []*ListNode{{Val: 4}, {Val: 5}, {Val: 1}, {Val: 9}},
				idx:  1,
			},
			output{
				ans: []int{4, 1, 9},
			},
		},
		{
			input{
				head: []*ListNode{{Val: 4}, {Val: 5}, {Val: 1}, {Val: 9}},
				idx:  2,
			},
			output{
				ans: []int{4, 5, 9},
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.head)-1; i++ {
			data.i.head[i].Next = data.i.head[i+1]
		}

		deleteNode(data.i.head[data.i.idx])
		head := data.i.head[0]
		res := []int{}
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		ast.Equal(data.o.ans, res)
	}
}
