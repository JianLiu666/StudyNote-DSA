package p00648

import "strings"

// Time Complexity: O(mn), where m is the number of dictionary, n is the length of each word
// Space Complexity: O(mn)
func replaceWords(dictionary []string, sentence string) string {
	trie := CreateTrie()
	for _, str := range dictionary {
		trie.Put(str)
	}

	strs := strings.Split(sentence, " ")
	var res strings.Builder
	for idx, str := range strs {
		res.WriteString(trie.Shortest(str))
		if idx != len(strs)-1 {
			res.WriteString(" ")
		}
	}

	return res.String()
}

type Node struct {
	Childs [26]*Node
	IsEnd  bool
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

// Time Complexity: O(n), where n is the length of string
// Space Complexity: O(n)
func (t *Trie) Put(str string) {
	current := t.Root
	for _, ch := range str {
		if current.Childs[ch-'a'] == nil {
			current.Childs[ch-'a'] = CreateNode()
		}
		current = current.Childs[ch-'a']
	}
	current.IsEnd = true
}

// Time Complexity: O(n), where n is the length of string
// Space Complexity: O(n)
func (t *Trie) Shortest(str string) string {
	current := t.Root

	var res strings.Builder
	for _, ch := range str {
		if current.IsEnd {
			return res.String()
		}
		if current.Childs[ch-'a'] == nil {
			return str
		}
		res.WriteRune(ch)
		current = current.Childs[ch-'a']
	}

	return str
}
