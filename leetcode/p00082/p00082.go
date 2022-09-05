package image

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 製造一個空的節點在 head 之前, 好方便後面作業
	dummy := &ListNode{Val: math.MinInt, Next: head}
	current := dummy

	for current != nil && current.Next != nil {
		// 每一次都從當前位置的下一個與下下一個設定指針檢查是否重複
		n1 := current.Next
		n2 := n1.Next

		// 控制 n2 一直往前尋訪直到兩個指針的 value 不一樣時才停止
		for n2 != nil && n1.Val == n2.Val {
			n2 = n2.Next
		}

		// 如果指針的相對位置不變, 代表 n1 不是重複數字, 所以可以直接位移
		// 反之 n1 包含重複元素, 我們要直接跳過 n1 這個數字, 直接指向 n2
		if n1.Next == n2 {
			current = current.Next
		} else {
			current.Next = n2
		}
	}

	return dummy.Next
}
