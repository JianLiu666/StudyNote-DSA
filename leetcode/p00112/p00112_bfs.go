package p00112

// Time Complexity: O(n)
// Space Complexity: O(n)
func hasPathSum_bfs(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	q := CreateQueue()
	q.Put(root, targetSum)
	for !q.Empty() {
		node, sum := q.Pop()
		sum -= node.Val
		if sum == 0 && node.Left == nil && node.Right == nil {
			return true
		}
		if node.Left != nil {
			q.Put(node.Left, sum)
		}
		if node.Right != nil {
			q.Put(node.Right, sum)
		}
	}

	return false
}

type QueueNode struct {
	Node *TreeNode
	Sum  int
	Next *QueueNode
	Prev *QueueNode
}

func CreateQueueNode(node *TreeNode, sum int) *QueueNode {
	return &QueueNode{
		Node: node,
		Sum:  sum,
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

func (q *Queue) Put(node *TreeNode, sum int) {
	queueNode := CreateQueueNode(node, sum)
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

func (q *Queue) Pop() (*TreeNode, int) {
	if q.size == 0 {
		return nil, -1
	}

	queueNode := q.tail
	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.tail = q.tail.Prev
		q.tail.Next = nil
	}
	queueNode.Next = nil
	queueNode.Prev = nil
	q.size--

	return queueNode.Node, queueNode.Sum
}
