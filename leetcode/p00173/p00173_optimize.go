package p00173

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack Stack
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{
		stack: CreateStack(),
	}
	iterator.helper(root)
	return iterator
}

func (this *BSTIterator) Next() int {
	node := this.stack.Pop()
	this.helper(node.Right)
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return !this.stack.Empty()
}

func (this *BSTIterator) helper(node *TreeNode) {
	for node != nil {
		this.stack.Put(node)
		node = node.Left
	}
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
