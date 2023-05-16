package p00024

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}

	cursor := dummy
	for cursor.Next != nil && cursor.Next.Next != nil {
		tmp := cursor.Next
		cursor.Next = tmp.Next
		tmp.Next = cursor.Next.Next
		cursor.Next.Next = tmp
		cursor = tmp
	}

	return dummy.Next
}
