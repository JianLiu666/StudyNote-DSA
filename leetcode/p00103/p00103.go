package p00103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(2^n), where n is the depth of tree
// Space Complexity: O(n)
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := CreateQueue()
	q.Put(root)

	res := [][]int{}

	level := 1
	for q.Size() > 0 {
		size := q.Size()
		vals := make([]int, size)

		for i := 0; i < size; i++ {
			node := q.Pop()
			if node.Left != nil {
				q.Put(node.Left)
			}
			if node.Right != nil {
				q.Put(node.Right)
			}

			if level&1 == 1 {
				vals[i] = node.Val
			} else {
				vals[size-1-i] = node.Val
			}
		}
		res = append(res, vals)
		level++
	}

	return res
}

type QueueNode struct {
	Val  *TreeNode
	Next *QueueNode
	Prev *QueueNode
}

func CreateQueueNode(val *TreeNode) *QueueNode {
	return &QueueNode{
		Val:  val,
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

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Put(val *TreeNode) {
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
