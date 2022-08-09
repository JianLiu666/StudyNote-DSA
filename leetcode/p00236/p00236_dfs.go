package p00236

// Time Complexity: O(n)
// Space Complexity: O(n)
func lowestCommonAncestor_dfs(root, p, q *TreeNode) *TreeNode {
	s := CreateStack()
	s.Put(root)

	visited := map[*TreeNode]bool{}
	memo := map[*TreeNode]bool{}

	for !s.Empty() {
		node := s.Pop()

		if (node.Left != nil && !visited[node.Left]) || (node.Right != nil && !visited[node.Right]) {
			s.Put(node)
			if node.Right != nil {
				s.Put(node.Right)
				visited[node.Right] = true
			}
			if node.Left != nil {
				s.Put(node.Left)
				visited[node.Left] = true
			}
		} else {
			mid := node == p || node == q
			left := false
			if node.Left != nil {
				left = memo[node.Left]
			}
			right := false
			if node.Right != nil {
				right = memo[node.Right]
			}

			memo[node] = mid || left || right

			if (mid && (left || right)) || (left && right) {
				return node
			}
		}
	}

	return nil
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
