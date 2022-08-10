package p00098

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func isValidBST(root *TreeNode) bool {
	res := []int{}

	s := CreateStrack()
	s.Put(root)

	memo := map[*TreeNode]bool{}
	for !s.Empty() {
		node := s.Pop()
		if node.Left == nil || memo[node.Left] {
			res = append(res, node.Val)
		}
		if !memo[node] {
			memo[node] = true
			if node.Right != nil {
				s.Put(node.Right)
			}
			if node.Left != nil {
				s.Put(node)
				s.Put(node.Left)
			}
		}
	}

	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}

	return true
}

type Stack struct {
	arr  []*TreeNode
	size int
}

func CreateStrack() Stack {
	return Stack{
		arr:  make([]*TreeNode, 0),
		size: 0,
	}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Put(val *TreeNode) {
	s.arr = append(s.arr, val)
	s.size++
}

func (s *Stack) Pop() *TreeNode {
	node := s.arr[s.size-1]
	s.size--
	s.arr = s.arr[:s.size]

	return node
}
