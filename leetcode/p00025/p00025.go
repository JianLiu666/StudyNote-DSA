package p00025

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	pre := dummy
	end := dummy

	for end != nil {
		// 將 end 定位到要翻轉的終點
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		// 暫存節點
		tmp := end.Next

		// 設定這次 group 開始翻轉的起點與終點
		start := pre.Next
		end.Next = nil

		// 翻轉這個 group, 並把前後的節點重新串回到正確的節點上
		pre.Next = reverse(start)
		start.Next = tmp

		// 將指針移動到下一個 group 的定位點
		pre = start
		end = start
	}

	return dummy.Next
}

// 翻轉 Linked List
// @param head node
//
// @return *ListNode 新的 head node
func reverse(head *ListNode) *ListNode {
	cur := head.Next
	pre := head
	pre.Next = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		cur, pre = tmp, cur
	}
	return pre
}
