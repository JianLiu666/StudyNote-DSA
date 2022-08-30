package p00211

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

type WordDictionary struct {
	Root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		Root: CreateNode(),
	}
}

func (this *WordDictionary) AddWord(word string) {
	current := this.Root
	for _, ch := range word {
		if current.Childs[ch-'a'] == nil {
			current.Childs[ch-'a'] = CreateNode()
		}
		current = current.Childs[ch-'a']
	}
	current.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(this.Root, word, 0)
}

func (this *WordDictionary) dfs(root *Node, word string, idx int) bool {
	if idx == len(word) {
		return root.IsEnd
	}

	if word[idx] != '.' && root.Childs[word[idx]-'a'] == nil {
		return false
	}

	if word[idx] != '.' {
		return this.dfs(root.Childs[word[idx]-'a'], word, idx+1)
	} else {
		for _, node := range root.Childs {
			if node == nil {
				continue
			}
			if found := this.dfs(node, word, idx+1); found {
				return true
			}
		}
	}

	return false
}
