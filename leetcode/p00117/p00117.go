package p00117

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	q := CreateQueue()
	putChildNodes(&q, root)

	for q.Size() > 0 {
		size := q.Size()
		left := q.Pop()
		putChildNodes(&q, left)

		for i := 1; i < size; i++ {
			right := q.Pop()
			left.Next = right
			putChildNodes(&q, right)
			left = right
		}
	}

	return root
}

func putChildNodes(q *Queue, node *Node) {
	if node.Left != nil {
		q.Put(node.Left)
	}
	if node.Right != nil {
		q.Put(node.Right)
	}
}

type QueueNode struct {
	Val  *Node
	Prev *QueueNode
	Next *QueueNode
}

func CreateQueueNode(val *Node) *QueueNode {
	return &QueueNode{
		Val:  val,
		Prev: nil,
		Next: nil,
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

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Put(val *Node) {
	node := CreateQueueNode(val)
	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		node.Next = q.head
		q.head.Prev = node
		q.head = node
	}
	q.size++
}

func (q *Queue) Pop() *Node {
	if q.size == 0 {
		return nil
	}

	node := q.tail
	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.tail = q.tail.Prev
		q.tail.Next = nil
	}
	node.Prev = nil
	node.Next = nil
	q.size--

	return node.Val
}
