package p00104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func maxDepth_bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := CreateQueue()
	q.Put(root)

	depth := 0
	for !q.Empty() {
		size := q.Size()
		depth++

		for i := 0; i < size; i++ {
			node := q.Pop()
			if node.Left != nil {
				q.Put(node.Left)
			}
			if node.Right != nil {
				q.Put(node.Right)
			}
		}
	}

	return depth
}

type QueueNode struct {
	Val  *TreeNode
	Next *QueueNode
	Prev *QueueNode
}

func CreateQueueNode(node *TreeNode) *QueueNode {
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

func (q *Queue) Put(treeNode *TreeNode) {
	queueNode := CreateQueueNode(treeNode)
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

func (q *Queue) Pop() *TreeNode {
	if q.size == 0 {
		return nil
	}

	queueNode := q.tail
	q.size--
	if q.size == 0 {
		q.head = nil
		q.tail = nil
	} else {
		q.tail = q.tail.Prev
		q.tail.Next = nil
	}
	queueNode.Next = nil
	queueNode.Prev = nil

	return queueNode.Val
}
