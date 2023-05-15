package p01721

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func swapNodes(head *ListNode, k int) *ListNode {
	left := head   // 準備 swapping 的左節點
	right := head  // 準備 swapping 的右節點
	cursor := head // 用來定位右節點位置的節點

	// 先找到準備 swapping 的左節點與準備用來定位右節點的 cursor node
	for i := 1; i < k; i++ {
		left = left.Next
		cursor = cursor.Next
	}

	// 同時移動 cursor, right 直到 curosr 走到 list 盡頭
	for cursor.Next != nil {
		right = right.Next
		cursor = cursor.Next
	}

	// swap
	left.Val, right.Val = right.Val, left.Val

	return head
}
