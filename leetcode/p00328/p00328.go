package p00328

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd, even := head, head.Next
	evenBegin := even

	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenBegin

	return head
}
