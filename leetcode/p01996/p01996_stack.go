package p01996

import "sort"

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func numberOfWeakCharacters_stack(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] < properties[j][0] {
			return true
		}
		if properties[i][0] == properties[j][0] {
			return properties[i][1] > properties[j][1]
		}
		return false
	})

	count := 0
	stack := []int{}
	for _, p := range properties {
		// attack 已經由小到大排序, 因此現在只需考慮如何處理 defense 即可
		for len(stack) != 0 && p[1] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			count++
		}
		stack = append(stack, p[1])
	}

	return count
}

type Stack struct {
	items []int // stored property index
	size  int
}

func CreateStack() Stack {
	return Stack{
		items: make([]int, 0),
		size:  0,
	}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Put(val int) {
	s.items = append(s.items, val)
	s.size++
}

func (s *Stack) Peek() int {
	return s.items[s.size-1]
}

func (s *Stack) Pop() int {
	val := s.items[s.size-1]
	s.size--
	s.items = s.items[:s.size]
	return val
}
