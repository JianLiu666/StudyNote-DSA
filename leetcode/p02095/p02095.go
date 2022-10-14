package p02095

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func deleteMiddle(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	slow := dummy
	fast := dummy

	step := 0
	for {
		step++

		fast = fast.Next
		if fast == nil {
			break
		}

		// 代表 fast pointer 移動速度是 slow pointer 的兩倍
		if step&1 == 0 {
			slow = slow.Next
		}
	}

	// 移除 middle node
	slow.Next = slow.Next.Next

	return dummy.Next
}
