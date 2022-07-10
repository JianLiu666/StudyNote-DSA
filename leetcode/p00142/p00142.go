package p00142

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func detectCycle_pointers(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			slow2 := head
			for slow != slow2 {
				slow = slow.Next
				slow2 = slow2.Next
			}
			return slow
		}
	}

	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func detectCycle_hashmap(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	memo := map[*ListNode]int{}

	current := head
	for current.Next != nil {
		if _, exists := memo[current]; exists {
			return current
		}
		memo[current] = 1
		current = current.Next
	}

	return nil
}
