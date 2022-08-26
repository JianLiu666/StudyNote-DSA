package p00677

type Node struct {
	childs [26]*Node
	sum    int
}

func CreateNode() *Node {
	return &Node{
		childs: [26]*Node{},
		sum:    0,
	}
}

type MapSum struct {
	root *Node
	memo map[string]int
}

func Constructor() MapSum {
	return MapSum{
		root: CreateNode(),
		memo: map[string]int{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	val, this.memo[key] = val-this.memo[key], val

	current := this.root
	for _, ch := range key {
		if current.childs[ch-97] == nil {
			current.childs[ch-97] = CreateNode()
		}
		current = current.childs[ch-97]
		current.sum += val
	}
}

func (this *MapSum) Sum(prefix string) int {
	current := this.root
	for _, ch := range prefix {
		if current.childs[ch-97] == nil {
			return 0
		}
		current = current.childs[ch-97]
	}
	return current.sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
