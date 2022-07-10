package p00141

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nodes []*ListNode
	pos   int
}

type output struct {
	ans bool
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
				ans: true,
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
				ans: true,
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
				ans: false,
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
				ans: false,
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.nodes)-1; i++ {
			data.i.nodes[i].Next = data.i.nodes[i+1]
		}
		if data.i.pos != -1 {
			data.i.nodes[len(data.i.nodes)-1].Next = data.i.nodes[data.i.pos]
		}

		ast.Equal(data.o.ans, hasCycle(data.i.nodes[0]))
	}
}
