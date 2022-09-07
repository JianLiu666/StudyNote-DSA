package p00230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(2^n)
// Space Complexity: O(n)
func kthSmallest_inorder(root *TreeNode, k int) int {
	s := CreateStack()
	s.Put(root)

	arr := []int{}
	for root != nil || !s.Empty() {
		for root != nil {
			s.Put(root)
			root = root.Left
		}
		root = s.Pop()
		arr = append(arr, root.Val)
		root = root.Right
	}

	return arr[k-1]
}

type Stack struct {
	items []*TreeNode
	size  int
}

func CreateStack() Stack {
	return Stack{
		items: make([]*TreeNode, 0),
		size:  0,
	}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Put(val *TreeNode) {
	s.items = append(s.items, val)
	s.size++
}

func (s *Stack) Pop() *TreeNode {
	val := s.items[s.size-1]
	s.size--
	s.items = s.items[:s.size]
	return val
}
