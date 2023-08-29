package p02462

import (
	"container/heap"
)

func totalCost(costs []int, k int, candidates int) int64 {
	result := 0

	h := &MinHeap{}
	heap.Init(h)

	left, right := 0, len(costs)-1

	// 根據 candidates 的規則, 從頭尾開始取出放進 min heap
	for left <= right && candidates > 0 {
		heap.Push(h, [2]int{costs[left], left})
		left++
		if left <= right {
			heap.Push(h, [2]int{costs[right], right})
			right--
		}
		candidates--
	}

	// 在 workers 可選人數還超過 candidates 時, 仍需要繼續考慮下一輪能放進去的 worker 人選
	for left <= right && k > 0 {
		mid := left + (right-left)/2
		worker := heap.Pop(h).([2]int)
		if worker[1] <= mid {
			heap.Push(h, [2]int{costs[left], left})
			left++
		} else {
			heap.Push(h, [2]int{costs[right], right})
			right--
		}

		result += worker[0]
		k--
	}

	// 補足剩下的人選
	for k > 0 {
		result += heap.Pop(h).([2]int)[0]
		k--
	}

	return int64(result)
}

type MinHeap [][2]int

func (h MinHeap) Len() int      { return len(h) }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool {
	if h[i][0] < h[j][0] {
		return true
	}
	if h[i][0] > h[j][0] {
		return false
	}
	return h[i][1] < h[j][1]
}
func (h *MinHeap) Push(v any) { *h = append(*h, v.([2]int)) }
func (h *MinHeap) Pop() (v any) {
	v, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]
	return
}
