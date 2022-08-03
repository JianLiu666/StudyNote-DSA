package p00024

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	current := head
	newHead := head.Next
	current.Next = current.Next.Next
	newHead.Next = current

	for current.Next != nil && current.Next.Next != nil {
		tmp := current.Next
		current.Next = current.Next.Next
		current = current.Next
		tmp.Next = current.Next
		current.Next = tmp
		current = current.Next
	}

	return newHead
}
