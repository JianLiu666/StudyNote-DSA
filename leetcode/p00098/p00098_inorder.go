package p00098

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func isValidBST_inorder(root *TreeNode) bool {
	res := []int{}

	s := CreateStrack()
	for root != nil || !s.Empty() {
		for root != nil {
			s.Put(root)
			root = root.Left
		}

		root = s.Pop()
		res = append(res, root.Val)
		root = root.Right
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
