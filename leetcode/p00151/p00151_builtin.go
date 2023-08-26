package p00151

import "strings"

func reverseWords_builtin(s string) string {
	words := strings.Fields(s)

	head, tail := 0, len(words)-1
	for head < tail {
		words[head], words[tail] = words[tail], words[head]
		head++
		tail--
	}

	return strings.Join(words, " ")
}
