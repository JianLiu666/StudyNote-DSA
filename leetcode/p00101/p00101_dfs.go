package p00101

// Time Complexity: O(n)
// Space Complexity: O(n)
func isSymmetric_dfs(root *TreeNode) bool {
	s := CreateStack()
	s.Put(root.Left)
	s.Put(root.Right)

	for !s.Empty() {
		right := s.Pop()
		left := s.Pop()

		if left != nil && right != nil {
			if left.Val != right.Val {
				return false
			}
			s.Put(left.Left)
			s.Put(right.Right)
			s.Put(right.Left)
			s.Put(left.Right)
		} else {
			if left != nil || right != nil {
				return false
			}
		}
	}

	return true
}

type Stack struct {
	arr  []*TreeNode
	size int
}

func CreateStack() Stack {
	return Stack{
		arr:  make([]*TreeNode, 0),
		size: 0,
	}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Put(node *TreeNode) {
	s.arr = append(s.arr, node)
	s.size++
}

func (s *Stack) Pop() *TreeNode {
	if s.size == 0 {
		return nil
	}

	node := s.arr[s.size-1]
	s.arr = s.arr[:s.size-1]
	s.size--

	return node
}
