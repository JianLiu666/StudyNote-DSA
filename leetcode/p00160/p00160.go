package p00160

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func getIntersectionNode_alignment(headA, headB *ListNode) *ListNode {
	currentA := headA
	currentB := headB
	lenA := 0
	lenB := 0

	for currentA != nil {
		lenA++
		currentA = currentA.Next
	}
	for currentB != nil {
		lenB++
		currentB = currentB.Next
	}

	currentA = headA
	for i := 0; i < lenA-lenB; i++ {
		currentA = currentA.Next
	}

	currentB = headB
	for i := 0; i < lenB-lenA; i++ {
		currentB = currentB.Next
	}

	for currentA != nil {
		if currentA == currentB {
			return currentA
		}

		currentA = currentA.Next
		currentB = currentB.Next
	}

	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func getIntersectionNode_cycle(headA, headB *ListNode) *ListNode {
	currentA := headA
	currentB := headB

	for currentA != currentB {
		if currentA == nil {
			currentA = headB
		} else {
			currentA = currentA.Next
		}

		if currentB == nil {
			currentB = headA
		} else {
			currentB = currentB.Next
		}
	}

	return currentA
}
