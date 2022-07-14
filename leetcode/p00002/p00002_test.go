package p00002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	l1 []*ListNode
	l2 []*ListNode
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
				l1: []*ListNode{{Val: 2}, {Val: 4}, {Val: 3}},
				l2: []*ListNode{{Val: 5}, {Val: 6}, {Val: 4}},
			},
			output{
				ans: []int{7, 0, 8},
			},
		},
		{
			input{
				l1: []*ListNode{{Val: 0}},
				l2: []*ListNode{{Val: 0}},
			},
			output{
				ans: []int{0},
			},
		},
		{
			input{
				l1: []*ListNode{{Val: 9}, {Val: 9}, {Val: 9}, {Val: 9}, {Val: 9}, {Val: 9}, {Val: 9}},
				l2: []*ListNode{{Val: 9}, {Val: 9}, {Val: 9}, {Val: 9}},
			},
			output{
				ans: []int{8, 9, 9, 9, 0, 0, 0, 1},
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.l1)-1; i++ {
			data.i.l1[i].Next = data.i.l1[i+1]
		}
		for i := 0; i < len(data.i.l2)-1; i++ {
			data.i.l2[i].Next = data.i.l2[i+1]
		}

		head := addTwoNumbers(data.i.l1[0], data.i.l2[0])
		res := []int{}
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		ast.Equal(data.o.ans, res)
	}
}
