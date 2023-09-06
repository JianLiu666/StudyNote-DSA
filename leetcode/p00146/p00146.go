package p00146

import "container/list"

type LRUCache struct {
	list     *list.List
	lookup   map[int]*list.Element
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list:     list.New(),
		lookup:   make(map[int]*list.Element),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, exists := this.lookup[key]; !exists {
		return -1
	}

	e := this.lookup[key]
	this.list.MoveToBack(e)

	return e.Value.([]int)[1]
}

func (this *LRUCache) Put(key int, value int) {
	if e, exists := this.lookup[key]; exists {
		e.Value.([]int)[1] = value
		this.list.MoveToBack(e)
	} else {
		newone := this.list.PushBack([]int{key, value})
		this.lookup[key] = newone
	}

	if this.list.Len() > this.capacity {
		delete(this.lookup, this.list.Front().Value.([]int)[0])
		this.list.Remove(this.list.Front())
	}
}

/**
["LRUCache","put", "put", "get","put", "get","put", "get","get","get"]
[[2],        [1,1], [2,2], [1],  [3,3], [2],  [4,4], [1],  [3],  [4]]
[null,       null,  null,  1,    null,  -1,   null,  -1,   3,    4] expected
[null,       null,  null,  1,    null,   2,   null,   1,   3,    4] mine

分析問題:
get: 只要檢查在或不在, 必須是 O(1) 操作, 用一個 hash table 維護 key: key, value: list node reference
put: 如果不存在就直接新增, 如果存在就更新, 一旦 size 超過 capacity 後要把最久沒有使用到的 key 刪除, 需要一個 linked list 來維護 [old, ..., new]
     每次刪除時都把 head pop, 一旦已存在的 node 更新就直接 append to tail
*/
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
