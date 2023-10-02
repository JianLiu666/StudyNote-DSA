package p01851

import (
	"container/heap"
	"sort"
)

func minInterval(intervals [][]int, queries []int) []int {
	result := make([]int, len(queries))

	pairs := [][2]int{}
	for i, n := range queries {
		pairs = append(pairs, [2]int{i, n})
		result[i] = -1
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &MinHeap{}
	i := 0
	for _, p := range pairs {
		idx, val := p[0], p[1]
		for i < len(intervals) && intervals[i][0] <= val {
			heap.Push(h, Window{
				size: intervals[i][1] - intervals[i][0] + 1,
				end:  intervals[i][1],
			})
			i++
		}

		for h.Len() > 0 && (*h)[0].end < val {
			heap.Pop(h)
		}

		if h.Len() > 0 {
			result[idx] = (*h)[0].size
		}
	}

	return result
}

type Window struct {
	size int
	end  int
}

type MinHeap []Window

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i].size < h[j].size }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(Window)) }
func (h *MinHeap) Pop() (v any) {
	v, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]
	return
}

/**
分析問題:
我們需要找到每個 queries[i] 在 intervals 裡面的對應到的最小區間
最直觀的做法就是 O(nm), n=len(intervals), m=len(queries)
嘗試降低每個 queries[i] 都要遍歷一次遍歷 intervals 的複雜度, 為此我們需要先把 intervals 跟 queries 的關係都確認下來, 所以:
 - 對 intervals 的 start time 做 ascending order 的排序
 - 對 queries 做 ascending order 排序

目前我們可以確定 queries[i] 是否落在 intervals[j] (至少 intervals[j][0] <= queries[i])
但是我們不能確保 queries[i] 是否也同時 <= intervals[j][1]

所以我們可以用一個 MinHeap 來解決這個問題, MinHeap 的 element 為 {interval range, end time}
且 comparison function 以 interval range 為主, 越小的在越前面, 這麼做的用意為:
 - 一旦確定 MinHeap 內的所有元素都涵蓋 queries[i] 時, 第一個一定是最小的區間

因此, 我們可以用 end time 來排空 MinHeap, 至此我們可以確保 queries[i] 一定落在 MinHeap 內, 因為:
 - 只有 intervals[j][0] <= queries[i] 的才會加入 MinHeap
 - 只要 intervals[j][1] < queries[i] 的都會被移出 MinHeap
*/
