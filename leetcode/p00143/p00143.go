package p00143

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func reorderList(head *ListNode) {
	stack := []*ListNode{}

	// step.1 find the middle of linked list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// step.2 push nodes of last halfs to the stack
	for slow != nil {
		stack = append(stack, slow)
		slow = slow.Next
	}

	// step.3 merge two half from first and last halfs together
	current := head
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node.Next = current.Next
		current.Next = node
		current = current.Next.Next
	}
	current.Next = nil
}
