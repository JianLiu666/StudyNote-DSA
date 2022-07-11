package p00019

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	second := head

	for i := 0; i < n; i++ {
		first = first.Next
	}

	for first != nil && first.Next != nil {
		first = first.Next
		second = second.Next
	}

	if first == nil && second == head {
		head = head.Next
	} else {
		second.Next = second.Next.Next
	}

	return head
}
