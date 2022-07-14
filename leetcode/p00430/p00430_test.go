package p00430

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	list1 := []*Node{{Val: 1}, {Val: 2}, {Val: 3}, {Val: 4}, {Val: 5}, {Val: 6}}
	list1[0].Next = list1[1]
	list1[len(list1)-1].Prev = list1[len(list1)-2]
	for i := 1; i < len(list1)-1; i++ {
		list1[i].Next = list1[i+1]
		list1[i].Prev = list1[i-1]
	}

	list2 := []*Node{{Val: 7}, {Val: 8}, {Val: 9}, {Val: 10}}
	list2[0].Next = list2[1]
	list2[len(list2)-1].Prev = list2[len(list2)-2]
	for i := 1; i < len(list2)-1; i++ {
		list2[i].Next = list2[i+1]
		list2[i].Prev = list2[i-1]
	}

	list1[2].Child = list2[0]

	list3 := []*Node{{Val: 11}, {Val: 12}}
	list3[0].Next = list3[1]
	list3[1].Prev = list3[0]

	list2[1].Child = list3[0]

	ast.Equal([]int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6}, checkDoublyLinkedList(flatten(list1[0])))

}

func checkDoublyLinkedList(root *Node) []int {
	dummy := &Node{Val: 0}
	dummy.Next = root
	current := root

	res := []int{}
	for current != nil {
		if current.Child != nil {
			return []int{}
		}
		if current != root && current.Prev != dummy {
			return []int{}
		}

		res = append(res, current.Val)
		current = current.Next
		dummy = dummy.Next
	}

	return res
}
