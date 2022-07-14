package p00430

type Node struct {
	Val   int
	Next  *Node
	Prev  *Node
	Child *Node
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func flatten(root *Node) *Node {
	recursive(root)
	return root
}

func recursive(root *Node) *Node {
	dummy := &Node{Val: 0}
	dummy.Next = root
	current := root

	for current != nil {
		if current.Child == nil {
			current = current.Next
			dummy = dummy.Next
		} else {
			tmp := current.Next
			child_tail := recursive(current.Child)
			child_tail.Next = current.Next
			if current.Next != nil {
				current.Next.Prev = child_tail
			}
			current.Next = current.Child
			current.Child.Prev = current
			current.Child = nil
			current = tmp
			dummy = child_tail
		}
	}

	return dummy
}
