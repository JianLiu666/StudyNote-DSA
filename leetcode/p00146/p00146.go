package p00146

import "container/list"

type LRUCache struct {
	list     *list.List
	data     map[int]*list.Element
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list:     list.New(),
		data:     make(map[int]*list.Element, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, exists := this.data[key]; !exists {
		return -1
	}

	// 保持最新的在前面
	elem := this.data[key]
	this.list.MoveToFront(elem)
	return elem.Value.([]int)[1]

}

func (this *LRUCache) Put(key int, value int) {
	// update operation
	if elem, exists := this.data[key]; exists {
		this.list.Remove(elem)
		newelem := this.list.PushFront([]int{key, value})
		this.data[key] = newelem
		return
	}

	// insert operation
	if len(this.data) == this.capacity {
		elem := this.list.Back()
		v := this.list.Remove(elem)
		delete(this.data, v.([]int)[0])
	}
	newelem := this.list.PushFront([]int{key, value})
	this.data[key] = newelem
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
