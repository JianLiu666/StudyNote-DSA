package bst

import (
	"strconv"
	"strings"
)

type BSTArray struct {
	bst []int
}

func CreateBSTArray(nums []int) BSTArray {
	instance := BSTArray{
		bst: make([]int, len(nums)),
	}
	for _, n := range nums {
		instance.Add(n)
	}

	return instance
}

func (t *BSTArray) String() string {
	var s strings.Builder
	if t.bst[0] == 0 {
		return s.String()
	}

	s.WriteString(strconv.Itoa(t.bst[0]))
	for i := 1; i < len(t.bst); i++ {
		s.WriteString(",")
		if t.bst[i] == 0 {
			s.WriteString("x")
		} else {
			s.WriteString(strconv.Itoa(t.bst[i]))
		}
	}

	return s.String()
}

func (t *BSTArray) Add(val int) {
	t.add(0, val)
}

func (t *BSTArray) add(nodeIdx, val int) {
	if nodeIdx >= len(t.bst) {
		t.expandSpace()
	}

	if t.bst[nodeIdx] == 0 {
		// found the leaf node to add
		t.bst[nodeIdx] = val
		return
	}

	if val == t.bst[nodeIdx] {
		// do something

	} else if val < t.bst[nodeIdx] {
		// move to left child node
		leftNodeIdx := (nodeIdx+1)*2 - 1
		t.add(leftNodeIdx, val)

	} else {
		// move to right child node
		rightNodeIdx := (nodeIdx + 1) * 2
		t.add(rightNodeIdx, val)
	}
}

func (t *BSTArray) Search(val int) int {
	return t.search(0, val)
}

func (t *BSTArray) search(nodeIdx, val int) (result int) {
	if nodeIdx >= len(t.bst) || t.bst[nodeIdx] == 0 {
		return
	}
	if t.bst[nodeIdx] == val {
		return val
	}

	if val < t.bst[nodeIdx] {
		leftNodeIdx := (nodeIdx+1)*2 - 1
		result = t.search(leftNodeIdx, val)
	} else {
		rightNodeIdx := (nodeIdx + 1) * 2
		result = t.search(rightNodeIdx, val)
	}

	return
}

func (t *BSTArray) expandSpace() {
	newBST := make([]int, len(t.bst)*2)
	for i, val := range t.bst {
		newBST[i] = val
	}
	t.bst = newBST
}
