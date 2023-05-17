package p02130

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	head []*ListNode
}

type output struct {
	ans int
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
				head: []*ListNode{{Val: 5}, {Val: 4}, {Val: 2}, {Val: 1}},
			},
			output{
				ans: 6,
			},
		},
		{
			input{
				head: []*ListNode{{Val: 4}, {Val: 2}, {Val: 2}, {Val: 3}},
			},
			output{
				ans: 7,
			},
		},
		{
			input{
				head: []*ListNode{{Val: 1}, {Val: 100000}},
			},
			output{
				ans: 100001,
			},
		},
	}

	for _, data := range tds {
		for i := 0; i < len(data.i.head)-1; i++ {
			data.i.head[i].Next = data.i.head[i+1]
		}

		// 不會異動到 linked list
		ast.Equal(data.o.ans, pairSum_sc_on(data.i.head[0]))

		// 會異動到 linked list, 懶得 deep copy data
		ast.Equal(data.o.ans, pairSum_sc_o1(data.i.head[0]))
	}
}
