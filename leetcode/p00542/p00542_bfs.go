package p00542

import "math"

// Time Complexity: O(mn), where m is the length of rows, and n is the length of columns
// Space Complexity: O(mn), where m is the length of rows, and n is the length of columns
func updateMatrix_bfs(mat [][]int) [][]int {
	res := make([][]int, len(mat))

	// initialization
	for row := 0; row < len(res); row++ {
		res[row] = make([]int, len(mat[0]))
		for col := 0; col < len(mat[0]); col++ {
			res[row][col] = math.MaxInt
		}
	}

	q := CreateQueue()

	// find start point
	for row, columns := range mat {
		for col, val := range columns {
			if val == 0 {
				res[row][col] = 0
				q.Put(row, col)
			}
		}
	}

	// calc distance
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for !q.Empty() {
		row, col := q.Pop()

		for i := 0; i < 4; i++ {
			newRow, newCol := row+directions[i][0], col+directions[i][1]

			// optimization without using continue syntax
			if newRow >= 0 && newRow < len(mat) &&
				newCol >= 0 && newCol < len(mat[0]) &&
				res[newRow][newCol] > res[row][col]+1 {

				res[newRow][newCol] = res[row][col] + 1
				q.Put(newRow, newCol)
			}
		}
	}

	return res
}

type node struct {
	Row  int
	Col  int
	Next *node
	Prev *node
}

func CreateNode(row, col int) *node {
	return &node{
		Row:  row,
		Col:  col,
		Next: nil,
		Prev: nil,
	}
}

type queue struct {
	head *node
	tail *node
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
	node := CreateNode(row, col)
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
	if q.size == 0 {
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
