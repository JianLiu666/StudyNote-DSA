package p00145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	s := CreateStack()
	s.Put(root)

	memo := map[*TreeNode]bool{}
	for !s.Empty() {
		node := s.Pop()

		if (node.Left == nil || memo[node.Left]) && (node.Right == nil || memo[node.Right]) {
			memo[node] = true
			res = append(res, node.Val)
		} else {
			s.Put(node)
			if node.Left != nil && !memo[node.Left] {
				s.Put(node.Left)
			} else if node.Right != nil && !memo[node.Right] {
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
