package p02130

// Time Complexity: O(n)
// Space Complexity: O(1)
func pairSum_sc_o1(head *ListNode) int {
	maximum := 0 // twin maximum

	dummy := &ListNode{Val: 0, Next: head}

	// step.1 找到 linked list 的中間位置
	slow := dummy
	fast := dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// step.2 反轉 linked list 右半部
	left := dummy.Next
	right := &ListNode{Val: 0, Next: nil}
	for slow != nil {
		tmp := slow.Next
		slow.Next = right
		right = slow
		slow = tmp
	}

	// step.3 compare and update maximum
	for left != nil && right != nil {
		if left.Val+right.Val > maximum {
			maximum = left.Val + right.Val
		}

		left = left.Next
		right = right.Next
	}

	return maximum
}
