package p00733

// Time Complexity: O(n)
// Space Complexity: O(n)
func floodFill_bfs(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	target := image[sr][sc]

	q := CreateQueue()
	q.Put(sr, sc)
	image[sr][sc] = color

	for !q.Empty() {
		size := q.Size()
		for i := 0; i < size; i++ {
			row, col := q.Pop()
			if row-1 >= 0 && image[row-1][col] == target {
				q.Put(row-1, col)
				image[row-1][col] = color
			}
			if row+1 < len(image) && image[row+1][col] == target {
				q.Put(row+1, col)
				image[row+1][col] = color
			}
			if col-1 >= 0 && image[row][col-1] == target {
				q.Put(row, col-1)
				image[row][col-1] = color
			}
			if col+1 < len(image[0]) && image[row][col+1] == target {
				q.Put(row, col+1)
				image[row][col+1] = color
			}
		}
	}

	return image
}

type queueNode struct {
	Row  int
	Col  int
	Next *queueNode
	Prev *queueNode
}

func CreateQueueNode(row, col int) *queueNode {
	return &queueNode{
		Row:  row,
		Col:  col,
		Next: nil,
		Prev: nil,
	}
}

type queue struct {
	head *queueNode
	tail *queueNode
	size int
}

func CreateQueue() queue {
	return queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *queue) Empty() bool {
	return q.size == 0
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) Put(row, col int) {
	node := CreateQueueNode(row, col)
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

func (q *queue) Pop() (int, int) {
	if q.Empty() {
		return -1, -1
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

	return node.Row, node.Col
}
