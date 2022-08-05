package p00144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	s := CreateStack()
	s.Put(root)

	for !s.Empty() {
		node := s.Pop()
		res = append(res, node.Val)
		if node.Right != nil {
			s.Put(node.Right)
		}
		if node.Left != nil {
			s.Put(node.Left)
		}
	}

	return res
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
