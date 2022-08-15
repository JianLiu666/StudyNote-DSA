package p00429

type Node struct {
	Val      int
	Children []*Node
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	q := CreateQueue()
	q.Put(root)

	for !q.Empty() {
		size := q.Size()
		tmp := []int{}
		for i := 0; i < size; i++ {
			node := q.Pop()
			tmp = append(tmp, node.Val)
			for _, child := range node.Children {
				q.Put(child)
			}
		}
		res = append(res, tmp)
	}

	return res
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
