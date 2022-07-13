package p00234

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		if fast == nil {
			tmp := slow.Next
			slow.Next = nil
			slow = tmp
		} else {
			slow = slow.Next
		}
	}

	previous := slow
	current := previous.Next
	previous.Next = nil

	for current != nil {
		tmp := current.Next
		current.Next = previous
		previous = current
		current = tmp
	}

	tail := previous
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}

		head = head.Next
		tail = tail.Next
	}

	return true
}
