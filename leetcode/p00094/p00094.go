package p00094

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	s := CreateStack()
	s.Put(root)

	visited := map[*TreeNode]bool{}
	for !s.IsEmpty() {
		node := s.Pop()

		if node.Left != nil && !visited[node.Left] {
			s.Put(node)
			s.Put(node.Left)
		} else {
			res = append(res, node.Val)
			visited[node] = true
			if node.Right != nil && !visited[node.Right] {
				s.Put(node.Right)
			}
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
