package p00234

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func isPalindrome(head *ListNode) bool {
	// step.1 find the middle of linked list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 當 linked list 長度是奇數時, 需要往後位移一個 node 才是要判斷回文的起點(終點)
	if fast != nil {
		slow = slow.Next
	}

	// step.2 reverse the half of linked list
	var last, current, next *ListNode
	current = slow
	for current != nil {
		next = current.Next
		current.Next = last
		last, current = current, next
	}

	// step.3 check palindrome
	left, right := head, last
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
