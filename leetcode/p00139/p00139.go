package p00139

// Time Complexity: O(n^2), where n is the length of s
// Space Complexity: O(nm), where n is the length of wordDic, and m is the length of wordDic[i]
func wordBreak(s string, wordDict []string) bool {
	t := CreateTrie()
	for _, word := range wordDict {
		t.insert(word)
	}

	return t.breakdown(s, 0)
}

type Node struct {
	children [26]*Node
	end      bool
}

type Trie struct {
	root   *Node
	pruned [301]bool // 題目給定的 s 範圍在 [1, 300], 為了提早剪枝
}

func CreateTrie() Trie {
	return Trie{
		root:   &Node{},
		pruned: [301]bool{},
	}
}

func (t *Trie) insert(word string) {
	cursor := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if cursor.children[idx] == nil {
			cursor.children[idx] = &Node{}
		}
		cursor = cursor.children[idx]
	}
	cursor.end = true
}

func (t *Trie) breakdown(word string, cur int) bool {
	if word[cur:] == "" {
		return true
	}
	// 先確認是否曾經從 cur 這個位置出發過, 如果早就失敗的話就不需要在重來一次了
	if t.pruned[cur] {
		return false
	}

	cursor := t.root

	for i := cur; i < len(word); i++ {
		idx := word[i] - 'a'
		if cursor.children[idx] == nil {
			// 如果從 word[cur] 這個位置開始查找的結果是失敗的話, 那就把位置記下來
			t.pruned[cur] = true
			return false
		}
		cursor = cursor.children[idx]
		if cursor.end && t.breakdown(word, i+1) {
			return true
		}
	}

	return cursor.end
}
