package p00135

import "container/heap"

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func candy(ratings []int) int {
	result := 0

	// 為了省略判斷邊界狀況，前後都補上一個
	ratings = append([]int{0}, ratings...)
	ratings = append(ratings, 0)
	ratings[0] = -1
	ratings[len(ratings)-1] = -1

	h := &MinHeap{}
	heap.Init(h)

	for i := 1; i < len(ratings)-1; i++ {
		heap.Push(h, &Child{
			rating: ratings[i],
			index:  i,
		})
	}

	// 用來記錄那些已經發過糖果的小孩都領到多少顆糖果了
	dp := make([]int, len(ratings))

	for h.Len() > 0 {
		child := heap.Pop(h).(*Child)
		if ratings[child.index-1] < child.rating && ratings[child.index+1] < child.rating {
			dp[child.index] = max(dp[child.index-1], dp[child.index+1]) + 1
		} else if ratings[child.index-1] >= child.rating && ratings[child.index+1] >= child.rating {
			dp[child.index] = 1
		} else {
			if ratings[child.index-1] < child.rating {
				dp[child.index] = dp[child.index-1] + 1
			} else {
				dp[child.index] = dp[child.index+1] + 1
			}
		}
	}

	for _, n := range dp {
		result += n
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Child struct {
	rating int
	index  int
}

type MinHeap []*Child

func (h MinHeap) Len() int      { return len(h) }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool {
	if h[i].rating < h[j].rating {
		return true
	}
	if h[i].rating > h[j].rating {
		return false
	}
	return h[i].index < h[j].index
}
func (h *MinHeap) Push(v any) { *h = append(*h, v.(*Child)) }
func (h *MinHeap) Pop() (v any) {
	v, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]
	return
}

/**
分析問題: 每個小孩至少一個, 且 rating 高的小孩要拿到比旁邊的人還多的糖果
用 minheap 從 rating 最少的孩子開始一個一個拿出來檢查
情境分析

[ v . v ] max(left, right) + 1
[ v . = ] or [ = . v ] or [ v . ^ ] or [ ^ . v ] -> v + 1
[ = . = ] or [ = . ^ ] or [ ^ . = ] or [ ^ . ^ ] -> 1

*/
