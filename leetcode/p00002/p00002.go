package p00002

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	current := dummy

	carried := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carried += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carried += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: carried % 10}
		current = current.Next
		carried /= 10
	}

	if carried > 0 {
		current.Next = &ListNode{Val: carried}
	}

	return dummy.Next
}
