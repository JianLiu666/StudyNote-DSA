package p00160

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	listA      []*ListNode
	listB      []*ListNode
	Intersects []*ListNode
}

type output struct {
	ans *ListNode
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
				listA:      []*ListNode{{Val: 4}, {Val: 1}},
				listB:      []*ListNode{{Val: 5}, {Val: 6}, {Val: 1}},
				Intersects: []*ListNode{{Val: 8}, {Val: 4}, {Val: 5}},
			},
			output{
				ans: nil,
			},
		},
		{
			input{
				listA:      []*ListNode{{Val: 1}, {Val: 9}, {Val: 1}},
				listB:      []*ListNode{{Val: 3}},
				Intersects: []*ListNode{{Val: 2}, {Val: 4}},
			},
			output{
				ans: nil,
			},
		},
		{
			input{
				listA:      []*ListNode{{Val: 2}, {Val: 6}, {Val: 4}},
				listB:      []*ListNode{{Val: 1}, {Val: 5}},
				Intersects: []*ListNode{},
			},
			output{
				ans: nil,
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.listA)-1; i++ {
			data.i.listA[i].Next = data.i.listA[i+1]
		}
		for i := 0; i < len(data.i.listB)-1; i++ {
			data.i.listB[i].Next = data.i.listB[i+1]
		}
		if len(data.i.Intersects) != 0 {
			data.i.listA[len(data.i.listA)-1].Next = data.i.Intersects[0]
			data.i.listB[len(data.i.listB)-1].Next = data.i.Intersects[0]
			data.o.ans = data.i.Intersects[0]
		}

		ast.Equal(data.o.ans, getIntersectionNode_alignment(data.i.listA[0], data.i.listB[0]))
		ast.Equal(data.o.ans, getIntersectionNode_cycle(data.i.listA[0], data.i.listB[0]))
	}
}
