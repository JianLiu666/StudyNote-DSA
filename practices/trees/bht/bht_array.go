package bht

import (
	"strconv"
	"strings"
)

// Using Array to implement Binary Heap Tree
type BHTArray struct {
	bht    []int
	endIdx int
}

func CreateBHTArray(nums []int) BHTArray {
	instance := BHTArray{
		bht:    make([]int, len(nums)),
		endIdx: len(nums) - 1,
	}
	for i, n := range nums {
		instance.bht[i] = n
	}

	// from BHT bottom to top
	for i := len(nums) - 1; i >= 0; i-- {
		instance.shiftDown(i)
	}

	return instance
}

func (t *BHTArray) String() string {
	var s strings.Builder
	if t.bht[0] == 0 {
		return s.String()
	}

	s.WriteString(strconv.Itoa(t.bht[0]))
	for i := 1; i <= t.endIdx; i++ {
		s.WriteString(",")
		if t.bht[i] == 0 {
			s.WriteString("x")
		} else {
			s.WriteString(strconv.Itoa(t.bht[i]))
		}
	}

	return s.String()
}

func (t *BHTArray) Remove() {
	t.bht[0], t.bht[t.endIdx] = t.bht[t.endIdx], t.bht[0]
	t.endIdx--

	t.shiftDown(0)
}

func (t *BHTArray) Add(val int) {
	if t.endIdx+1 > len(t.bht)-1 {
		t.expandSpace()
	}

	t.endIdx++
	t.bht[t.endIdx] = val
	t.shiftUp(t.endIdx)
}

func (t *BHTArray) shiftDown(idx int) {
	if idx > t.endIdx {
		return
	}

	childIdx := t.getBiggerChildIdx(idx)
	if childIdx == -1 {
		return
	}

	if t.bht[idx] < t.bht[childIdx] {
		t.bht[idx], t.bht[childIdx] = t.bht[childIdx], t.bht[idx]
		t.shiftDown(childIdx)
	}
}

func (t *BHTArray) shiftUp(idx int) {
	if idx <= 0 {
		return
	}

	parantIdx := (idx+1)/2 - 1
	if t.bht[idx] > t.bht[parantIdx] {
		t.bht[idx], t.bht[parantIdx] = t.bht[parantIdx], t.bht[idx]
		t.shiftUp(parantIdx)
	}
}

func (t *BHTArray) getBiggerChildIdx(idx int) (childIdx int) {
	childIdx = -1

	lIdx := (idx+1)*2 - 1
	rIdx := (idx + 1) * 2

	if lIdx <= t.endIdx && rIdx <= t.endIdx {
		if t.bht[lIdx] >= t.bht[rIdx] {
			childIdx = lIdx
		} else {
			childIdx = rIdx
		}
	} else if lIdx <= t.endIdx {
		childIdx = lIdx
	} else if rIdx <= t.endIdx {
		childIdx = rIdx
	}

	return
}

func (t *BHTArray) expandSpace() {
	newBHT := make([]int, len(t.bht)*2)
	for i, val := range t.bht {
		newBHT[i] = val
	}
	t.bht = newBHT
}
