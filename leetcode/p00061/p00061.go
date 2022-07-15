package p00061

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	size := 0
	current := head
	for current != nil {
		size++
		current = current.Next
	}

	k %= size

	tail := head
	for i := 0; i < k; i++ {
		tail = tail.Next
	}

	anchor := head
	for tail.Next != nil {
		tail = tail.Next
		anchor = anchor.Next
	}

	tail.Next = head
	new_head := anchor.Next
	anchor.Next = nil

	return new_head
}
