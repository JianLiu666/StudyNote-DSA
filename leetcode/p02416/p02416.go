package p02416

// Time Complexity: O(nm), where n is the number of words, m is the length of word
// Space Complexity: O(m)
func sumPrefixScores(words []string) []int {
	trie := CreateTrie()
	for _, word := range words {
		trie.Add(word)
	}

	res := make([]int, len(words))
	for i, word := range words {
		res[i] = trie.Acc(word)
	}
	return res
}

type TrieNode struct {
	childs [26]*TrieNode
	acc    int
}

func CreateTrieNode() *TrieNode {
	return &TrieNode{
		childs: [26]*TrieNode{},
		acc:    0,
	}
}

type Trie struct {
	root *TrieNode
}

func CreateTrie() Trie {
	return Trie{
		root: CreateTrieNode(),
	}
}

func (t *Trie) Add(word string) {
	current := t.root
	for _, ch := range word {
		if current.childs[ch-'a'] == nil {
			current.childs[ch-'a'] = CreateTrieNode()
		}
		current = current.childs[ch-'a']
		current.acc++
	}
}

func (t *Trie) Acc(word string) int {
	current := t.root
	acc := 0
	for _, ch := range word {
		if current.childs[ch-'a'] == nil {
			return acc
		}
		current = current.childs[ch-'a']
		acc += current.acc
	}
	return acc
}
