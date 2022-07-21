package p00092

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	var reverse_head *ListNode
	var reverse_head_prev *ListNode
	var reverse_tail *ListNode

	var previous *ListNode
	current := head

	step := 1
	for current != nil {
		if step == left {
			reverse_head_prev = previous
			reverse_tail = current

			tmp := current.Next
			current.Next = nil
			previous = current
			current = tmp

		} else if step == right {
			reverse_tail.Next = current.Next
			current.Next = previous

			if reverse_head_prev != nil {
				reverse_head_prev.Next = current
				return head
			} else {
				reverse_head = current
				return reverse_head
			}

		} else if step > left && step < right {
			tmp := current.Next
			current.Next = previous
			previous = current
			current = tmp

		} else {
			previous = current
			current = current.Next
		}

		step++
	}

	return nil
}
