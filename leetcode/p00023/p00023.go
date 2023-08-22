package p00023

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)

	for _, root := range lists {
		for root != nil {
			heap.Push(h, root.Val)
			root = root.Next
		}
	}

	dummy := &ListNode{}
	root := dummy
	for h.Len() > 0 {
		root.Next = &ListNode{Val: heap.Pop(h).(int)}
		root = root.Next
	}

	return dummy.Next
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
