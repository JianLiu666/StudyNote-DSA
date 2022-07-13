package p00021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	list1 []*ListNode
	list2 []*ListNode
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
				list1: []*ListNode{{Val: 1}, {Val: 2}, {Val: 4}},
				list2: []*ListNode{{Val: 1}, {Val: 3}, {Val: 4}},
			},
			output{
				ans: []int{1, 1, 2, 3, 4, 4},
			},
		},
		{
			input{
				list1: []*ListNode{nil},
				list2: []*ListNode{nil},
			},
			output{
				ans: []int{},
			},
		},
		{
			input{
				list1: []*ListNode{nil},
				list2: []*ListNode{{Val: 0}},
			},
			output{
				ans: []int{0},
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.list1)-1; i++ {
			data.i.list1[i].Next = data.i.list1[i+1]
		}
		for i := 0; i < len(data.i.list2)-1; i++ {
			data.i.list2[i].Next = data.i.list2[i+1]
		}

		head := mergeTwoLists(data.i.list1[0], data.i.list2[0])
		res := []int{}
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		ast.Equal(data.o.ans, res)
	}
}
