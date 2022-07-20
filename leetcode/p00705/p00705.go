package p00705

type HashSet_BruteForce struct {
	arr []bool
}

func Constructor() HashSet_BruteForce {
	return HashSet_BruteForce{
		arr: make([]bool, 1000001),
	}
}

func (this *HashSet_BruteForce) Add(key int) {
	this.arr[key] = true
}

func (this *HashSet_BruteForce) Remove(key int) {
	this.arr[key] = false
}

func (this *HashSet_BruteForce) Contains(key int) bool {
	return this.arr[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
