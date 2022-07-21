package p00705

type MyHashSet struct {
	buckets [][]int
}

func hasher(key int) int {
	return ((key * 786433) & ((1 << 20) - 1)) >> 7
}

func contains(arr []int, key int) int {
	for i, num := range arr {
		if num == key {
			return i
		}
	}

	return -1
}

func Constructor() MyHashSet {
	return MyHashSet{
		buckets: make([][]int, 1<<13),
	}
}

func (this *MyHashSet) Add(key int) {
	hashed := hasher(key)
	if contains(this.buckets[hashed], key) == -1 {
		this.buckets[hashed] = append(this.buckets[hashed], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	hashed := hasher(key)
	if idx := contains(this.buckets[hashed], key); idx > -1 {
		this.buckets[hashed] = append(this.buckets[hashed][:idx], this.buckets[hashed][idx+1:]...)
	}
}

func (this *MyHashSet) Contains(key int) bool {
	if contains(this.buckets[hasher(key)], key) == -1 {
		return false
	}

	return true
}
