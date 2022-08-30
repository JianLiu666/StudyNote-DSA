package p00212

// Time Complexity: O(m*n*4^10),
//  - where
//     m is the height of board,
//     n is the width of board,
//     4^10 means each cell has 4 directions can traversal and the maximum of each word is 10
// Space Complexity: O(m*n), where m is the number of words, n is the length of each word
func findWords(board [][]byte, words []string) []string {
	// 紀錄行走過的路線
	visited := [13][13]bool{}
	// 紀錄符合的單字
	existed := map[string]bool{}

	trie := CreateTrie()
	for _, word := range words {
		trie.Put(word)
	}

	for row, columns := range board {
		for col, val := range columns {
			helper(board, visited, existed, &trie, string(val), row, col)
		}
	}

	res := []string{}
	for word := range existed {
		res = append(res, word)
	}
	return res
}

func helper(board [][]byte, visited [13][13]bool, existed map[string]bool, trie *Trie, word string, row, col int) {
	exist := trie.Exists(word)
	if exist == 0 {
		return
	}
	if exist == 2 && !existed[word] {
		existed[word] = true
		trie.Delete(word)
	}

	visited[row][col] = true
	if row-1 >= 0 && !visited[row-1][col] {
		helper(board, visited, existed, trie, word+string(board[row-1][col]), row-1, col)
	}
	if row+1 < len(board) && !visited[row+1][col] {
		helper(board, visited, existed, trie, word+string(board[row+1][col]), row+1, col)
	}
	if col-1 >= 0 && !visited[row][col-1] {
		helper(board, visited, existed, trie, word+string(board[row][col-1]), row, col-1)
	}
	if col+1 < len(board[0]) && !visited[row][col+1] {
		helper(board, visited, existed, trie, word+string(board[row][col+1]), row, col+1)
	}
	visited[row][col] = false
}

type Node struct {
	Childs    [26]*Node
	NumChilds int
	IsEnd     bool
}

func CreateNode() *Node {
	return &Node{
		Childs: [26]*Node{},
		IsEnd:  false,
	}
}

type Trie struct {
	Root *Node
}

func CreateTrie() Trie {
	return Trie{
		Root: CreateNode(),
	}
}

// Time Complexity: O(n), where n is the length of word
// Space Complexity: O(n)
func (t *Trie) Put(s string) {
	current := t.Root
	for _, ch := range s {
		if current.Childs[ch-'a'] == nil {
			current.Childs[ch-'a'] = CreateNode()
			current.NumChilds++
		}
		current = current.Childs[ch-'a']
	}
	current.IsEnd = true
}

// Time Complexity: O(n), where n is the length of word
// Space Complexity: O(n)
func (t *Trie) Delete(s string) {
	stack := []*Node{}
	current := t.Root
	for _, ch := range s {
		current = current.Childs[ch-'a']
		stack = append(stack, current)
	}

	current, stack = stack[len(stack)-1], stack[:len(stack)-1]
	if len(stack)+1 == len(s) && current.NumChilds > 0 {
		return
	}
	for len(stack) > 0 {
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if current.NumChilds > 1 {
			return
		}
		current.Childs[s[len(stack)]-'a'] = nil
	}
}

// 查找是否有符合的單字
// 0: 未命中, 1: 符合 prefix, 2: 完全命中
// Time Complexity: O(n), where n is the length of word
// Space Complexity: O(1)
func (t *Trie) Exists(s string) int {
	current := t.Root
	for _, ch := range s {
		if current.Childs[ch-'a'] == nil {
			return 0
		}
		current = current.Childs[ch-'a']
	}
	if current.IsEnd {
		return 2
	}
	return 1
}
