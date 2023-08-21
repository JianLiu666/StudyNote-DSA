package p00212

var Directions [4][2]int = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func findWords(board [][]byte, words []string) []string {
	result := []string{}

	trieRoot := &TrieNode{}
	visited := map[[2]int]bool{}

	// 把所有的字串都加進 trie
	for _, word := range words {
		cursor := trieRoot
		for _, ch := range word {
			idx := ch - 'a'
			if cursor.children[idx] == nil {
				cursor.children[idx] = &TrieNode{}
			}
			cursor = cursor.children[idx]
			cursor.count++
		}
		cursor.end = true
	}

	for row, columns := range board {
		for col := range columns {
			dfs(board, row, col, visited, trieRoot, trieRoot, []byte{}, &result)
		}
	}

	return result
}

func dfs(board [][]byte, row, col int, visited map[[2]int]bool, trieRoot, trieCursor *TrieNode, word []byte, result *[]string) {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) {
		return
	}
	if visited[[2]int{row, col}] {
		return
	}
	// 只要當前的 curosr 沒辦法繼續在 trie 裡面走下去, 那就可以提早結束了 (early pruning)
	if trieCursor.children[board[row][col]-'a'] == nil || trieCursor.children[board[row][col]-'a'].count == 0 {
		return
	}

	// 只有 trieCursor 不用退回的原因是因為雖然這是 reference
	// 但是 golang 在 function call 傳遞的參數永遠都是 call by value, 所以只是把這個地址在複製一份當作參數
	// 現在的這個 trieCursor 指向的地址不會因為 dfs 的過程被改變
	trieCursor = trieCursor.children[board[row][col]-'a']
	word = append(word, board[row][col])

	// 命中之後就把這個字串從 trie 裡面移除, 減少以後重複查找的必要, 也是 early pruning
	if trieCursor.end {
		*result = append(*result, string(word))
		removeWord(trieRoot, string(word))
	}

	visited[[2]int{row, col}] = true
	for _, offset := range Directions {
		nextRow, nextCol := row+offset[0], col+offset[1]
		dfs(board, nextRow, nextCol, visited, trieRoot, trieCursor, word, result)
	}
	visited[[2]int{row, col}] = false

	word = word[:len(word)-1]

	return
}

type TrieNode struct {
	children [26]*TrieNode
	end      bool
	count    int
}

func removeWord(root *TrieNode, word string) {
	for _, ch := range word {
		idx := ch - 'a'
		root = root.children[idx]
		root.count--
	}
	root.end = false
}
