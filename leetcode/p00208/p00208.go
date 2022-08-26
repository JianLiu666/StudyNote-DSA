package p00208

type Node struct {
	childs [26]*Node
	isEnd  bool
}

func CreateNode() *Node {
	return &Node{
		childs: [26]*Node{},
		isEnd:  false,
	}
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: CreateNode(),
	}
}

func (this *Trie) Insert(word string) {
	current := this.root
	for _, ch := range word {
		if current.childs[ch-97] == nil {
			current.childs[ch-97] = CreateNode()
		}
		current = current.childs[ch-97]
	}
	current.isEnd = true
}

func (this *Trie) Search(word string) bool {
	current := this.root
	for _, ch := range word {
		if current.childs[ch-97] == nil {
			return false
		}
		current = current.childs[ch-97]
	}
	return current.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.root
	for _, ch := range prefix {
		if current.childs[ch-97] == nil {
			return false
		}
		current = current.childs[ch-97]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
