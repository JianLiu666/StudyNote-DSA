package p02130

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func pairSum_sc_on(head *ListNode) int {
	maximum := 0   // twin maximum
	tmp := []int{} // 用來暫存 twin 中第一個數值

	dummy := &ListNode{Val: 0, Next: head}

	// step.1 找到 linked list 的中間位置
	slow := dummy
	fast := dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		tmp = append(tmp, slow.Val)

		fast = fast.Next.Next
	}

	// step.2 compare and update maximum
	for i := len(tmp) - 1; i >= 0; i-- {
		slow = slow.Next
		tmp[i] += slow.Val

		if tmp[i] > maximum {
			maximum = tmp[i]
		}
	}

	return maximum
}
