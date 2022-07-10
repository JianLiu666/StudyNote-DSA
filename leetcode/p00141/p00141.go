package p00141

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next

		if fast.Next == nil {
			return false
		}

		fast = fast.Next
		if slow.Next == fast.Next {
			return true
		}
	}

	return false
}
