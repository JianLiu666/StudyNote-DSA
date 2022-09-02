package p00637

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func averageOfLevels_bfs(root *TreeNode) []float64 {
	queue := CreateQueue()
	queue.Put(root)

	res := []float64{}
	for queue.Size() > 0 {
		size := queue.Size()
		sum := 0
		for i := 0; i < size; i++ {
			node := queue.Pop()
			sum += node.Val
			if node.Left != nil {
				queue.Put(node.Left)
			}
			if node.Right != nil {
				queue.Put(node.Right)
			}
		}
		res = append(res, float64(sum)/float64(size))
	}

	return res
}

type QueueNode struct {
	Val  *TreeNode
	Prev *QueueNode
	Next *QueueNode
}

func CreateQueueNode(val *TreeNode) *QueueNode {
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

func (q *Queue) Put(val *TreeNode) {
	node := CreateQueueNode(val)
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		node.Next = q.head
		q.head.Prev = node
		q.head = node
	}
	q.size++
}

func (q *Queue) Pop() *TreeNode {
	node := q.tail
	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.tail = q.tail.Prev
		q.tail.Next = nil
	}
	node.Prev = nil
	q.size--

	return node.Val
}
