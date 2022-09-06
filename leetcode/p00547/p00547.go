package p00547

// Time Complexity: O(n^2)
// Space Complexity: O(n^2)
func findCircleNum(isConnected [][]int) int {
	size := len(isConnected)
	u := CreateUnionFind(size)

	for i := 0; i < size; i++ {
		// 每個城市只要檢查一次就好不必重複檢查, pruning 技巧
		for j := i + 1; j < size; j++ {
			if isConnected[i][j] == 1 {
				u.Union(i, j)
			}
		}
	}

	return u.Count()
}

type UnionFind struct {
	parent []int
	rank   []int
	count  int // 記錄州的數量
}

func CreateUnionFind(size int) UnionFind {
	instance := UnionFind{
		parent: make([]int, size),
		rank:   make([]int, size),
		count:  size,
	}

	// 每一個節點預設的 parent/root 都是自己本身
	for i := 0; i < size; i++ {
		instance.parent[i] = i
	}

	return instance
}

func (u *UnionFind) Count() int {
	return u.count
}

// 找到子集合中的 root index
func (u *UnionFind) Find(p int) int {
	// 持續往上追蹤直到找到 p 子集合中的 root index 為止
	for p != u.parent[p] {
		// 每次查找時都把子集合的這條鏈路壓縮 1/2, 提高未來查找時的效率
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}

	return p
}

// 將兩個給定的子集合 p, q 關聯在一起
func (u *UnionFind) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	// 兩個給定的節點本身就屬於同一個子集合
	if rootP == rootQ {
		return
	}

	// 比較兩個子集合的權重, 將小子集合歸屬到大的子集合中
	if u.rank[rootQ] > u.rank[rootP] {
		u.parent[rootP] = u.parent[rootQ]
	} else {
		u.parent[rootQ] = u.parent[rootP]
		// 如果兩個子集合的權重相同時, 那就歸屬到 p 的子集合中
		if u.rank[rootP] == u.rank[rootQ] {
			u.rank[rootP]++
		}
	}

	// 將兩個城市合併到同一個州, 因此州的數量-1
	u.count--
}
