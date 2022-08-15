package p00559

type Node struct {
	Val      int
	Children []*Node
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	q := CreateQueue()
	q.Put(root)

	depth := 0
	for !q.Empty() {
		depth++
		size := q.Size()
		for i := 0; i < size; i++ {
			node := q.Pop()
			for _, child := range node.Children {
				q.Put(child)
			}
		}
	}

	return depth
}

type QueueNode struct {
	Val  *Node
	Next *QueueNode
	Prev *QueueNode
}

func CreateQueueNode(node *Node) *QueueNode {
	return &QueueNode{
		Val:  node,
		Next: nil,
		Prev: nil,
	}
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
	size int
}

func CreateQueue() Queue {
	return Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *Queue) Empty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Put(node *Node) {
	queueNode := CreateQueueNode(node)
	if q.size == 0 {
		q.head = queueNode
		q.tail = queueNode
	} else {
		queueNode.Next = q.head
		q.head.Prev = queueNode
		q.head = queueNode
	}
	q.size++
}

func (q *Queue) Pop() *Node {
	queueNode := q.tail
	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.tail = q.tail.Prev
		q.tail.Next = nil
	}
	queueNode.Prev = nil
	q.size--

	return queueNode.Val
}
