package p00695

// Time Complexity: O(mn), where m is the height of grid, n is the width of grid
// Space Complexity: O(mn)
func maxAreaOfIsland_DFS(grid [][]int) int {
	max_row, max_col := len(grid), len(grid[0])

	s := CreateStack()

	max_area := 0
	for row, columns := range grid {
		for col, val := range columns {
			if val != 1 {
				continue
			}
			s.Put(row, col)
			area := 0
			for !s.Empty() {
				r, c := s.Pop()
				if grid[r][c] != 1 {
					continue
				}

				grid[r][c] = 0
				area++

				if r-1 >= 0 && grid[r-1][c] == 1 {
					s.Put(r-1, c) // up
				}
				if r+1 < max_row && grid[r+1][c] == 1 {
					s.Put(r+1, c) // down
				}
				if c-1 >= 0 && grid[r][c-1] == 1 {
					s.Put(r, c-1) // left
				}
				if c+1 < max_col && grid[r][c+1] == 1 {
					s.Put(r, c+1) // right
				}
			}
			if area > max_area {
				max_area = area
			}
		}
	}

	return max_area
}

type Stack struct {
	arr  [][2]int
	size int
}

func CreateStack() Stack {
	return Stack{
		arr:  make([][2]int, 0),
		size: 0,
	}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Put(row, column int) {
	s.arr = append(s.arr, [2]int{row, column})
	s.size++
}

func (s *Stack) Pop() (int, int) {
	pos := s.arr[s.size-1]
	s.size--
	s.arr = s.arr[:s.size]
	return pos[0], pos[1]
}
