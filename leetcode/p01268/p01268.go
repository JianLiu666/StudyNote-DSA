package p01268

import "sort"

func suggestedProducts(products []string, searchWord string) [][]string {
	result := [][]string{}

	sort.Strings(products)

	trieRoot := &Node{}
	for i, product := range products {
		cursor := trieRoot
		for _, ch := range product {
			idx := ch - 'a'
			if cursor.children[idx] == nil {
				cursor.children[idx] = &Node{}
			}
			cursor = cursor.children[idx]
			cursor.backlog = append(cursor.backlog, i)
		}
		cursor.end = true
	}

	cursor := trieRoot
	for i, ch := range searchWord {
		idx := ch - 'a'
		if cursor.children[idx] == nil {
			cursor.children[idx] = &Node{}
		}
		cursor = cursor.children[idx]
		result = append(result, []string{})
		for j := 0; j < len(cursor.backlog) && len(result[i]) < 3; j++ {
			result[i] = append(result[i], products[cursor.backlog[j]])
		}
	}

	return result
}

type Node struct {
	children [26]*Node
	end      bool
	backlog  []int
}
