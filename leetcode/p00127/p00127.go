package p00127

func ladderLength(beginWord string, endWord string, wordList []string) int {
	result := 1

	graph := map[string][]string{}
	visited := map[string]map[string]bool{}

	// 轉換成 graph 儲存
	for i := 0; i < len(wordList); i++ {
		if _, exists := graph[wordList[i]]; !exists {
			graph[wordList[i]] = []string{}
			visited[wordList[i]] = map[string]bool{}
		}

		for j := i + 1; j < len(wordList); j++ {
			if _, exists := graph[wordList[j]]; !exists {
				graph[wordList[j]] = []string{}
				visited[wordList[j]] = map[string]bool{}
			}
			if canTransformate(wordList[i], wordList[j]) {
				graph[wordList[i]] = append(graph[wordList[i]], wordList[j])
				graph[wordList[j]] = append(graph[wordList[j]], wordList[i])
			}
		}
	}

	q := []string{}
	if list, exists := graph[beginWord]; exists {
		for _, word := range list {
			q = append(q, word)
		}
	} else {
		for _, word := range wordList {
			if canTransformate(beginWord, word) {
				q = append(q, word)
			}
		}
	}

	for len(q) > 0 {
		size := len(q)
		result++
		for ; size > 0; size-- {
			word := q[0]
			q = q[1:]

			if word == endWord {
				return result
			}

			for _, next := range graph[word] {
				if visited[word][next] {
					continue
				}
				visited[word][next] = true
				q = append(q, next)
			}
		}
	}

	return 0
}

func canTransformate(a, b string) bool {
	missed := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		missed++
		if missed > 1 {
			return false
		}
	}
	return true
}
