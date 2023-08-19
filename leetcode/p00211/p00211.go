package p00211

type Node struct {
	children [26]*Node
	end      bool
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	cursor := this.root
	for _, ch := range word {
		idx := ch - 'a'
		if cursor.children[idx] == nil {
			cursor.children[idx] = &Node{}
		}
		cursor = cursor.children[idx]
	}
	cursor.end = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(this.root, word, 0)
}

func (this *WordDictionary) dfs(root *Node, word string, idx int) bool {
	if idx == len(word) {
		return root.end
	}

	if word[idx] == '.' {
		for _, child := range root.children {
			if child == nil {
				continue
			}
			if this.dfs(child, word, idx+1) {
				return true
			}
		}
		return false
	}

	if root.children[word[idx]-'a'] == nil {
		return false
	}
	return this.dfs(root.children[word[idx]-'a'], word, idx+1)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
