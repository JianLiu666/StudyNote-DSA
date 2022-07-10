package p00142

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nodes []*ListNode
	pos   int
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
				nodes: []*ListNode{
					{Val: 3, Next: nil},
					{Val: 2, Next: nil},
					{Val: 0, Next: nil},
					{Val: -4, Next: nil},
				},
				pos: 1,
			},
			output{
				ans: nil,
			},
		},
		{
			input{
				nodes: []*ListNode{
					{Val: 1, Next: nil},
					{Val: 2, Next: nil},
				},
				pos: 0,
			},
			output{
				ans: nil,
			},
		},
		{
			input{
				nodes: []*ListNode{
					{Val: 1, Next: nil},
				},
				pos: -1,
			},
			output{
				ans: nil,
			},
		},
		{
			input{
				nodes: []*ListNode{
					{Val: 1, Next: nil},
					{Val: 1, Next: nil},
					{Val: 1, Next: nil},
					{Val: 1, Next: nil},
				},
				pos: -1,
			},
			output{
				ans: nil,
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.nodes)-1; i++ {
			data.i.nodes[i].Next = data.i.nodes[i+1]
		}
		if data.i.pos != -1 {
			data.i.nodes[len(data.i.nodes)-1].Next = data.i.nodes[data.i.pos]
			data.o.ans = data.i.nodes[data.i.pos]
		}

		ast.Equal(data.o.ans, detectCycle_pointers(data.i.nodes[0]))
		ast.Equal(data.o.ans, detectCycle_hashmap(data.i.nodes[0]))
	}
}
