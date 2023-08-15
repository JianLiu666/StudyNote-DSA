package p00148

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(nlogn), where n is the length of linked list
// Space Complexity: O(1)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// step.1 find the middle of linked list, and cut to two halfs
	slow, fast := head, head
	var middle *ListNode
	for fast != nil && fast.Next != nil {
		middle = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	middle.Next = nil

	// step.2 sort list each half
	l1 := sortList(head)
	l2 := sortList(slow)

	// step.3 merge lists and return
	return merge(l1, l2)

}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}
