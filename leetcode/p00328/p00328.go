package p00328

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	previous := head
	current := head.Next

	for current != nil && current.Next != nil {
		tmp := previous.Next
		previous.Next = current.Next
		current.Next = current.Next.Next
		previous = previous.Next
		previous.Next = tmp
		current = current.Next
	}

	return head
}
