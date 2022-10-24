package p00021

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cursor := dummy

	for list1 != nil || list2 != nil {
		tmp := &ListNode{}

		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				tmp.Val = list1.Val
				list1 = list1.Next
			} else {
				tmp.Val = list2.Val
				list2 = list2.Next
			}
		} else if list1 == nil {
			tmp.Val = list2.Val
			list2 = list2.Next
		} else {
			tmp.Val = list1.Val
			list1 = list1.Next
		}

		cursor.Next = tmp
		cursor = cursor.Next
	}

	return dummy.Next
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
