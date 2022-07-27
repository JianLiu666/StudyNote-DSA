package p00114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func flatten_stack(root *TreeNode) {
	if root == nil {
		return
	}

	s := CreateStack()
	dfs_stack(&s, root)

	current := s.Pop()
	for !s.IsEmpty() {
		current.Left = nil
		current.Right = s.Pop()
		current = current.Right
	}
}

func dfs_stack(s *Stack, node *TreeNode) {
	if node.Right != nil {
		dfs_stack(s, node.Right)
	}
	if node.Left != nil {
		dfs_stack(s, node.Left)
	}
	s.Put(node)
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

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Put(node *TreeNode) {
	s.arr = append(s.arr, node)
	s.size++
}

func (s *Stack) Pop() *TreeNode {
	if s.IsEmpty() {
		return nil
	}

	node := s.arr[s.size-1]
	s.arr = s.arr[:s.size-1]
	s.size--

	return node
}
