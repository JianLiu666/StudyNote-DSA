package p00394

import (
	"strconv"
	"strings"
)

// Time Complextiy: O(n)
// Space Complextiy: O(n)
func decodeString(s string) string {
	var result strings.Builder

	stack := CreateStack()

	head := 0
	status := 0
	for i := 0; i < len(s); i++ {
		// processing number
		if status == 0 && s[i] >= '0' && s[i] <= '9' {
			head = i
			status = 1
		}
		if status == 1 && (s[i] >= '0' && s[i] <= '9') && (s[i+1] < '0' || s[i+1] > '9') {
			stack.Put(s[head : i+1])
			status = 0
		}

		// processing alphabet
		if status == 0 && s[i] >= 'a' && s[i] <= 'z' {
			head = i
			status = 2
		}
		if status == 2 && (s[i] >= 'a' && s[i] <= 'z') && (i+1 == len(s) || (s[i+1] < 'a' || s[i+1] > 'z')) {
			if stack.Empty() {
				result.WriteString(s[head : i+1])
			} else {
				stack.Put(s[head : i+1])
			}
			status = 0
		}

		// processing right brackets
		if s[i] == ']' {
			letters := stack.Pop()
			for stack.Top()[0] >= 'a' && stack.Top()[0] <= 'z' {
				letters = stack.Pop() + letters
			}

			count, _ := strconv.Atoi(stack.Pop())
			var concated strings.Builder
			for i := 0; i < count; i++ {
				concated.WriteString(letters)
			}

			if stack.Empty() {
				result.WriteString(concated.String())
			} else {
				stack.Put(concated.String())
			}
		}
	}

	return result.String()
}

type Stack struct {
	strs []string
	size int
}

func CreateStack() Stack {
	return Stack{
		strs: make([]string, 0),
		size: 0,
	}
}

func (c *Stack) Size() int {
	return c.size
}

func (s *Stack) Empty() bool {
	return s.size == 0
}
func (s *Stack) Put(val string) {
	s.strs = append(s.strs, val)
	s.size++
}

func (s *Stack) Pop() string {
	val := s.strs[s.size-1]
	s.strs = s.strs[:s.size-1]
	s.size--

	return val
}

func (s *Stack) Top() string {
	return s.strs[s.size-1]
}
