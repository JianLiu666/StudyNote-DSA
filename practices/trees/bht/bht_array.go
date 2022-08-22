package bht

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
