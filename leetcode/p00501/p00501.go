package p00501

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n), where n is the number of nodes
// Space Complexity: O(m), where m is the depth of tree
func findMode(root *TreeNode) []int {
	global_maximum := 0
	local_maximum := 0
	prev := 100001

	s := CreateStack()

	// in-order traversal
	res := []int{}
	for root != nil || !s.Empty() {
		for root != nil {
			s.Put(root)
			root = root.Left
		}
		root = s.Pop()

		// main logic
		if prev != root.Val {
			local_maximum = 0
			prev = root.Val
		}

		local_maximum++
		if local_maximum == global_maximum {
			res = append(res, prev)
		} else if local_maximum > global_maximum {
			global_maximum = local_maximum
			res = []int{prev}
		}

		root = root.Right
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
	node := s.arr[s.size-1]
	s.size--
	s.arr = s.arr[:s.size]
	return node
}
