package p00206

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head.Next
	previous := head
	previous.Next = nil

	for current != nil {
		next := current.Next
		current.Next = previous
		current, previous = next, current
	}

	return previous
}
