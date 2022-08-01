package p00733

// Time Complexity: O(n)
// Space Complexity: O(n)
func floodFill_dfs(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	stack := CreateStack()
	stack.Put(sr, sc)

	target := image[sr][sc]
	for !stack.Empty() {
		row, col := stack.Pop()
		// image[row][col] = color

		if row-1 >= 0 && image[row-1][col] == target {
			stack.Put(row-1, col)
			image[row-1][col] = color
		}
		if row+1 < len(image) && image[row+1][col] == target {
			stack.Put(row+1, col)
			image[row+1][col] = color
		}
		if col-1 >= 0 && image[row][col-1] == target {
			stack.Put(row, col-1)
			image[row][col-1] = color
		}
		if col+1 < len(image[0]) && image[row][col+1] == target {
			stack.Put(row, col+1)
			image[row][col+1] = color
		}
	}

	return image
}

type stackNode struct {
	Row int
	Col int
}

func CreateStackNode(row, col int) *stackNode {
	return &stackNode{
		Row: row,
		Col: col,
	}
}

type stack struct {
	arr  []*stackNode
	size int
}

func CreateStack() stack {
	return stack{
		arr:  make([]*stackNode, 0),
		size: 0,
	}
}

func (s *stack) Empty() bool {
	return s.size == 0
}

func (s *stack) Put(row, col int) {
	s.arr = append(s.arr, CreateStackNode(row, col))
	s.size++
}

func (s *stack) Pop() (int, int) {
	node := s.arr[s.size-1]
	s.arr = s.arr[:s.size-1]
	s.size--

	return node.Row, node.Col
}
