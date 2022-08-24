package p00220

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	k    int
	t    int
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
			input{nums: []int{1, 2, 3, 1}, k: 3, t: 0},
			output{ans: true},
		},
		{
			input{nums: []int{1, 0, 1, 1}, k: 1, t: 2},
			output{ans: true},
		},
		{
			input{nums: []int{1, 5, 9, 1, 5, 9}, k: 2, t: 3},
			output{ans: false},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, containsNearbyAlmostDuplicate(data.i.nums, data.i.k, data.i.t), idx+1)
	}
}

func TestBuiltinSearh(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(sort.Search(len(arr), func(i int) bool {
		return arr[i] > 55
	}))

	arr = []int{}
	fmt.Println(sort.Search(len(arr), func(i int) bool {
		return arr[i] > 55
	}))

	arr = []int{10, 20, 30, 40}
	arr = arr[4:]
	fmt.Println(arr)
}
