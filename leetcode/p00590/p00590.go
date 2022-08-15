package p00590

type Node struct {
	Val      int
	Children []*Node
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func postorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}

	s := CreateStack()
	s.Put(root)
	for !s.Empty() {
		node := s.Pop()
		res = append(res, node.Val)
		for _, child := range node.Children {
			s.Put(child)
		}
	}

	head, tail := 0, len(res)-1
	for head < tail {
		res[head], res[tail] = res[tail], res[head]
		head++
		tail--
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
