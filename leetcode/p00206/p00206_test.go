package p00206

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
				ans: []int{5, 4, 3, 2, 1},
			},
		},
		{
			input{
				head: []*ListNode{{Val: 1}, {Val: 2}},
			},
			output{
				ans: []int{2, 1},
			},
		},
		{
			input{
				head: []*ListNode{nil},
			},
			output{
				ans: []int{},
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.head)-1; i++ {
			data.i.head[i].Next = data.i.head[i+1]
		}

		head := reverseList(data.i.head[0])
		res := []int{}
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		ast.Equal(data.o.ans, res)
	}
}
