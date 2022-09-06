package p00143

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func reorderList(head *ListNode) {
	stack := CreateStack()

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	for slow != nil {
		stack.Put(slow)
		slow = slow.Next
	}

	current := head
	for !stack.Empty() {
		node := stack.Pop()
		node.Next = current.Next
		current.Next = node
		current = current.Next.Next
	}
	current.Next = nil
}

type Stack struct {
	arr  []*ListNode
	size int
}

func CreateStack() Stack {
	return Stack{
		arr:  make([]*ListNode, 0),
		size: 0,
	}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Put(val *ListNode) {
	s.arr = append(s.arr, val)
	s.size++
}

func (s *Stack) Pop() *ListNode {
	val := s.arr[s.size-1]
	s.size--
	s.arr = s.arr[:s.size]
	return val
}
