package p00086

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func partition(head *ListNode, x int) *ListNode {
	left_head := &ListNode{}
	left_cursor := left_head

	right_head := &ListNode{}
	right_cursor := right_head

	for head != nil {
		if head.Val < x {
			left_cursor.Next = head
			left_cursor = left_cursor.Next
		} else {
			right_cursor.Next = head
			right_cursor = right_cursor.Next
		}
		head = head.Next
	}

	left_cursor.Next = right_head.Next
	right_cursor.Next = nil

	return left_head.Next
}
