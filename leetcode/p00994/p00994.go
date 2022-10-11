package p00994

// Time Complexity: O(mn), where m is the height of grid, n is the width of grid
// Space Complexity: O(mn)
func orangesRotting(grid [][]int) int {
	q := CreateQueue()

	// 遍歷一次找到所有新鮮/腐爛的水果
	fresh := 0
	for row, columns := range grid {
		for col, val := range columns {
			if val == 1 {
				fresh++
			} else if val == 2 {
				q.Put(row, col)
			}
		}
	}

	// BFS 搜尋方向
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 執行 BFS 直到所有能尋訪到的路徑都探索完畢
	times := 0
	for !q.Empty() {
		size := q.Size()
		rotting := false

		for i := 0; i < size; i++ {
			row, col := q.Pop()
			for _, pos := range directions {
				next_r := row + pos[0]
				next_c := col + pos[1]
				if next_r < 0 || next_r >= len(grid) || next_c < 0 || next_c >= len(grid[0]) || grid[next_r][next_c] != 1 {
					continue
				}
				// 更新狀態
				grid[next_r][next_c] = 2
				fresh--
				rotting = true
				q.Put(next_r, next_c)
			}
		}

		if rotting {
			times++
		}
	}

	if fresh > 0 {
		return -1
	}
	return times
}

type Node struct {
	Row  int
	Col  int
	Next *Node
	Prev *Node
}

func CreateNode(row, col int) *Node {
	return &Node{
		Row:  row,
		Col:  col,
		Next: nil,
		Prev: nil,
	}
}

type Queue struct {
	head *Node
	tail *Node
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

func (q *Queue) Put(row, col int) {
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

func (q *Queue) Pop() (row, col int) {
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
	q.size--

	return node.Row, node.Col
}
