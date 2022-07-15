package p00138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	memo := make(map[*Node]*Node, 0)

	anchor := head
	root := &Node{Val: head.Val}
	memo[anchor] = root

	anchor = anchor.Next
	current := root

	for anchor != nil {
		node := &Node{Val: anchor.Val}
		current.Next = node
		current = current.Next

		memo[anchor] = current

		anchor = anchor.Next
	}

	anchor = head
	current = root

	for anchor != nil {
		if node, exists := memo[anchor.Random]; exists {
			current.Random = node
		}

		anchor = anchor.Next
		current = current.Next
	}

	return root
}
