package p01249

import "strings"

// Time Complexity: O(n), where n is the length of s
// Space Complexity: O(n)
func minRemoveToMakeValid_index(s string) string {
	stack := CreateIndexStack()
	skip := map[int]bool{}

	for i, ch := range s {
		if ch == '(' {
			stack.Put(i)
		} else if ch == ')' {
			if stack.Empty() {
				skip[i] = true
			} else {
				stack.Pop()
			}
		}
	}

	for !stack.Empty() {
		skip[stack.Pop()] = true
	}

	var res strings.Builder
	for i, ch := range s {
		if !skip[i] {
			res.WriteRune(ch)
		}
	}

	return res.String()
}

type IndexStack struct {
	arr  []int
	size int
}

func CreateIndexStack() IndexStack {
	return IndexStack{
		arr:  []int{},
		size: 0,
	}
}

func (s *IndexStack) Empty() bool {
	return s.size == 0
}

func (s *IndexStack) Put(val int) {
	s.arr = append(s.arr, val)
	s.size++
}

func (s *IndexStack) Peek() int {
	return s.arr[s.size-1]
}

func (s *IndexStack) Pop() int {
	val := s.arr[s.size-1]
	s.size--
	s.arr = s.arr[:s.size]
	return val
}
