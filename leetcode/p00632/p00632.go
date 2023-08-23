package p00632

import (
	"container/heap"
	"math"
)

func smallestRange(nums [][]int) []int {
	// 用來記錄最後輸出時的左右端點元素
	left, right := 0, 0

	minRange := math.MaxInt

	// golang 不像 c++ 有好用的 set, 所以對於這道題目的最大值而言, 我們只能用一個變數去維護
	// 根據每次迭代的過程可以確保 maxVal 會逐步提高, 直到有任何一個 nums[i] 裡面已經沒有元素可用時
	// maxVal 就會是題目要的 smallest range 的最大值
	maxVal := math.MinInt

	h := &MinHeap{}
	heap.Init(h)

	// 把所有陣列的第一個元素都加進 min-heap 內
	for idx, list := range nums {
		heap.Push(h, &Element{
			num:    list[0],
			numIdx: 0,
			arrIdx: idx,
		})
		maxVal = max(maxVal, list[0])
	}

	for {
		e := h.Peek()

		// greedy
		if maxVal-e.num < minRange {
			minRange = maxVal - e.num
			left, right = e.num, maxVal
		}

		// pop 前先確認陣列還有元素可用來比較
		if e.numIdx+1 == len(nums[e.arrIdx]) {
			break
		}
		heap.Pop(h)
		heap.Push(h, &Element{
			num:    nums[e.arrIdx][e.numIdx+1],
			numIdx: e.numIdx + 1,
			arrIdx: e.arrIdx,
		})
		maxVal = max(maxVal, nums[e.arrIdx][e.numIdx+1])
	}

	return []int{left, right}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Element struct {
	num    int
	numIdx int // 紀錄元素在自己陣列裡面的索引
	arrIdx int // 記錄這個陣列在 nums 裡面的索引
}

type MinHeap []*Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Peek() *Element     { return h[0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i].num < h[j].num }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(*Element)) }
func (h *MinHeap) Pop() (v any) {
	v, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]
	return
}

/**

   4     10    15          24 26
 0     9    12       20
     5            18    22       30

4 10 15 24 26
0 9 12 20
5 18 22 30

0 4 5 -> 5
4 5 9 -> 5
5 9 10
9 10 18
10 12 18
12 15 18
15 18 20
18 20 24
20 22 24

*/
