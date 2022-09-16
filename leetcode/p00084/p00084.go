package p00084

// Time Complexity: O(n)
// Space Complexity: O(n)
func largestRectangleArea(heights []int) int {
	// 前後都補 0 的原因是為了讓後續實作 monotonic stack 時不需要額外考慮邊界判斷
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)

	maximum := 0

	s := CreateStack()
	for i, h := range heights {
		// 只要現在這個位置 heights[i] 的高度比 stack[-1] 還要小時, 就表示此時 stack[-1] 的高度是目前最指最高的高度
		// 我們可以也必須先將這個高度的所佔面積先計算掉, 確保後續不用再考慮這個凸點的面積
		for !s.Empty() && heights[s.Peek()] > h {
			// 先把這個高度 pop 出來
			cur := s.Pop()
			// 以 Example.1 為例, 當 index 指到 heights[4]=2 時, 表示當前為止最大的高度就是6, 要開始逐一往回計算過去的面積
			// 控制 left 指針來逐一計算這些過去存在 stack 內位置的個別面積
			left := s.Peek() + 1
			right := i - 1

			maximum = max(maximum, (right-left+1)*heights[cur])
		}
		s.Put(i)
	}

	return maximum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Stack struct {
	idxes []int
	size  int
}

func CreateStack() Stack {
	return Stack{
		idxes: make([]int, 0),
		size:  0,
	}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Put(idx int) {
	s.idxes = append(s.idxes, idx)
	s.size++
}

func (s *Stack) Peek() int {
	return s.idxes[s.size-1]
}

func (s *Stack) Pop() int {
	idx := s.idxes[s.size-1]
	s.size--
	s.idxes = s.idxes[:s.size]
	return idx
}
