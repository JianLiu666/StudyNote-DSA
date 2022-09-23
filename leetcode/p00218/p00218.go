package p00218

import (
	"container/heap"
	"sort"
)

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func getSkyline(buildings [][]int) [][]int {
	res := [][]int{}

	// 使用 golang package 實作 priority queue (max-heap)
	h := &NodeHeap{}
	heap.Init(h)

	// 將 buidings 的線段轉成 1D array (左, 右節點)
	edgePoints := []Node{}
	for i := 0; i < len(buildings); i++ {
		edgePoints = append(edgePoints, Node{
			x:      buildings[i][0],
			y:      buildings[i][2],
			idx:    i,
			isLeft: true,
			pos:    0,
		})
		edgePoints = append(edgePoints, Node{
			x:      buildings[i][1],
			y:      buildings[i][2],
			idx:    i,
			isLeft: false,
			pos:    0,
		})
	}

	// 重新排序邊緣節點
	// tc: O(nlogn)
	sort.Slice(edgePoints, func(i, j int) bool {
		if edgePoints[i].x != edgePoints[j].x {
			return edgePoints[i].x < edgePoints[j].x
		} else {
			return edgePoints[i].y < edgePoints[j].y
		}
	})

	// 為了讓右邊的水平線輪廓可以被畫出來
	heap.Push(h, &Node{
		x:      0,
		y:      0,
		idx:    0,
		isLeft: false,
		pos:    0,
	})

	visited := make([]*Node, len(buildings))
	for i := 0; i < len(edgePoints); i++ {
		p := edgePoints[i]

		if p.isLeft {
			// 如果是左節點的話, 就將新線段寫進 priority queue
			p.pos = h.Len()
			heap.Push(h, &p)
			visited[p.idx] = &p
		} else {
			// 反之則從 priority queue 移除該線段
			heap.Remove(h, visited[p.idx].pos)
		}

		// 只要 x 座標有變動就必須紀錄 cornor point, 才能描繪出建築輪廓
		if i == len(edgePoints)-1 || p.x != edgePoints[i+1].x {
			currMaxHeight := (*h)[0].y
			if len(res) == 0 || currMaxHeight != res[len(res)-1][1] {
				res = append(res, []int{p.x, currMaxHeight})
			}
		}
	}

	return res
}

type Node struct {
	x      int
	y      int
	idx    int
	isLeft bool
	pos    int
}

type NodeHeap []*Node

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].pos, h[j].pos = i, j
}

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].y > h[j].y
}

func (h *NodeHeap) Push(v interface{}) {
	*h = append(*h, v.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
