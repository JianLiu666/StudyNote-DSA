package p00203

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	head []*ListNode
	val  int
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
				head: []*ListNode{{Val: 1}, {Val: 2}, {Val: 6}, {Val: 3}, {Val: 4}, {Val: 5}, {Val: 6}},
				val:  6,
			},
			output{
				ans: []int{1, 2, 3, 4, 5},
			},
		},
		{
			input{
				head: []*ListNode{nil},
				val:  1,
			},
			output{
				ans: []int{},
			},
		},
		{
			input{
				head: []*ListNode{{Val: 7}, {Val: 7}, {Val: 7}, {Val: 7}},
				val:  7,
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

		head := removeElements(data.i.head[0], data.i.val)
		res := []int{}
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		ast.Equal(data.o.ans, res)
	}
}
