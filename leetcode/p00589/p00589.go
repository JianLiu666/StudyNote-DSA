package p00589

type Node struct {
	Val      int
	Children []*Node
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	s := CreateStack()
	s.Put(root)

	res := []int{}
	for !s.Empty() {
		node := s.Pop()
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			s.Put(node.Children[i])
		}
	}

	return res
}

type Stack struct {
	arr  []*Node
	size int
}

func CreateStack() Stack {
	return Stack{
		arr:  make([]*Node, 0),
		size: 0,
	}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Put(node *Node) {
	s.arr = append(s.arr, node)
	s.size++
}

func (s *Stack) Pop() *Node {
	if s.size == 0 {
		return nil
	}

	node := s.arr[s.size-1]
	s.size--
	s.arr = s.arr[:s.size]
	return node
}
