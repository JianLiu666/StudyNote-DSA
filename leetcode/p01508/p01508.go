package p01508

import "container/heap"

// Time Complexity: O(klogn), where k is the number of right, n is the number of nums
// Space Complexity: O(n)
func rangeSum(nums []int, n int, left int, right int) int {
	result := 0

	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			heap.Push(h, sum)
		}
	}

	for i := 1; i < left; i++ {
		heap.Pop(h)
	}
	for ; left <= right; left++ {
		result += heap.Pop(h).(int)
		result %= 1000000007
	}

	return result
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(int)) }
func (h *MinHeap) Pop() (v any) {
	v, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]
	return
}
