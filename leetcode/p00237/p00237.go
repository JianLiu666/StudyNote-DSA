package p00237

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
