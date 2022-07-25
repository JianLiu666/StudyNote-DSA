package p00200

// Time Complexity: O(mn)
// Space Complexity: O(n)
func numIslands(grid [][]byte) int {
	grid_height := len(grid) - 1
	grid_width := len(grid[0]) - 1

	// represent the left, top, right, down direction of the cursor
	checkList := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

	// SC: O(n)
	q := CreateQueue(len(grid) * len(grid[0]))
	numIslands := 0

	// traverse the whole grid -> TC: O(n)
	for row, columns := range grid {
		for col, val := range columns {
			if val == '1' {
				q.Put([]int{row, col})
				grid[row][col] = '2'

				// traverse this island -> TC: O(m)
				for !q.IsEmpty() {
					pos, _ := q.Pop()

					// TC: O(4)
					for _, offset := range checkList {
						y := pos[0] + offset[0]
						x := pos[1] + offset[1]
						if x >= 0 && x <= grid_width && y >= 0 && y <= grid_height &&
							grid[y][x] == '1' {
							q.Put([]int{y, x})
							grid[y][x] = '2'
						}
					}
				}

				numIslands++
			}
		}
	}

	return numIslands
}

type Queue struct {
	queue    [][]int
	head     int
	tail     int
	size     int
	capacity int
}

func CreateQueue(capacity int) Queue {
	return Queue{
		queue:    make([][]int, capacity),
		head:     0,
		tail:     capacity - 1,
		size:     0,
		capacity: capacity,
	}
}

func (this *Queue) IsEmpty() bool {
	return this.size == 0
}

func (this *Queue) IsFull() bool {
	return this.size == this.capacity
}

func (this *Queue) Put(position []int) bool {
	if this.IsFull() {
		return false
	}

	this.tail++
	if this.tail == this.capacity {
		this.tail = 0
	}

	this.queue[this.tail] = position
	this.size++

	return true
}

func (this *Queue) Pop() ([]int, bool) {
	if this.IsEmpty() {
		return nil, false
	}

	pos := this.queue[this.head]

	this.head++
	if this.head == this.capacity {
		this.head = 0
	}

	this.size--
	return pos, true
}
