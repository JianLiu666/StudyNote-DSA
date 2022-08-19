package p00083

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	for current != nil && current.Next != nil {
		tmp := current.Next
		for tmp != nil && tmp.Val == current.Val {
			tmp = tmp.Next
		}
		current.Next = tmp
		current = current.Next
	}

	return head
}
