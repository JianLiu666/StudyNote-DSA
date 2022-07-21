package p00706

type Node struct {
	Key int
	Val int
}
type MyHashMap struct {
	buckets [][]*Node
}

func hasher(key int) int {
	return ((key * 786433) & (1<<20 - 1)) >> 7
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets: make([][]*Node, 1<<13),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	hashed := hasher(key)

	for _, node := range this.buckets[hashed] {
		if node.Key == key {
			node.Val = value
			return
		}
	}
	this.buckets[hashed] = append(this.buckets[hashed], &Node{Key: key, Val: value})
}

func (this *MyHashMap) Get(key int) int {
	for _, node := range this.buckets[hasher(key)] {
		if node.Key == key {
			return node.Val
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	hashed := hasher(key)

	idx := -1
	for i, node := range this.buckets[hashed] {
		if node.Key == key {
			idx = i
			break
		}
	}

	if idx == -1 {
		return
	} else if idx == 0 {
		this.buckets[hashed] = append([]*Node{}, this.buckets[hashed][1:]...)
	} else if idx == len(this.buckets[hashed])-1 {
		this.buckets[hashed] = append([]*Node{}, this.buckets[hashed][:idx]...)
	} else {
		this.buckets[hashed] = append(this.buckets[hashed][:idx], this.buckets[hashed][idx+1:]...)
	}
}
