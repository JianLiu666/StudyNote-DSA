package p02542

import (
	"container/heap"
	"sort"
)

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func maxScore(nums1 []int, nums2 []int, k int) int64 {
	var res int = 0
	var sum int = 0

	// 先將兩個 slice 按照相同順序由大到小排序完畢
	slice := customSlice{
		nums1: nums1,
		nums2: nums2,
	}
	sort.Sort(slice)

	h := &minHeap{}
	heap.Init(h)

	// 開始從頭遍歷 nums2 的數組, 在遍歷的過程中不斷找出 global maximum
	for i, multiplier := range slice.nums2 {
		if h.Len() == k {
			// tc: O(logn)
			sum -= heap.Pop(h).(int)
		}

		sum += slice.nums1[i]
		// tc: O(logn)
		heap.Push(h, slice.nums1[i])

		if h.Len() == k {
			// 這裡可以直接用 multiplier 的原因是因為不管 nums1 pop 出哪一個 element 又加入哪一個 element
			// 當下的 multiplier 一定是最小的 (i.e. 這就是我們一開始把 nums2 由大到小排序的目的)
			res = max(res, sum*multiplier)
		}
	}

	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 自定一個 slice structure 來同時維護 nums1, nums2 的 index ordering
type customSlice struct {
	nums1 []int
	nums2 []int
}

func (s customSlice) Len() int {
	return len(s.nums2)
}

func (s customSlice) Less(i, j int) bool {
	// using '>' reverses the sort -- biggest first
	return s.nums2[i] > s.nums2[j]
}

func (s customSlice) Swap(i, j int) {
	s.nums1[i], s.nums1[j] = s.nums1[j], s.nums1[i]
	s.nums2[i], s.nums2[j] = s.nums2[j], s.nums2[i]
}

// 自定一個 min heap (priority queue) 來維護 k 個 nums1 的數組
type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {
	(*h) = append((*h), x.(int))
}

func (h *minHeap) Pop() any {
	// heap 的 pop 操作會先把 minimum element 放到 slice 的最尾端
	// 所以我們在這邊才可以直接拿最後一個 element, 同時順手修剪 slice 的長度
	i := (*h).Len() - 1
	res := (*h)[i]
	(*h) = (*h)[:i]
	return res
}
