package p00399

// Time Complexity: O(n^2), where n is the number of equations
// Space Complexity: O(n), where n is the number of equations
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	u := CreateUnionFind()

	// 建立 Union Find Set
	// tc: O(n)
	for i, equation := range equations {
		u.Add(equation[0])
		u.Add(equation[1])
		// tc: O(n)
		u.Union(equation[0], equation[1], values[i])
	}

	res := make([]float64, len(queries))
	for i, equation := range queries {
		res[i] = u.Query(equation[0], equation[1])
	}
	return res
}

type UnionFind struct {
	parent map[string]string
	weight map[string]float64
}

func CreateUnionFind() UnionFind {
	return UnionFind{
		parent: make(map[string]string, 0),
		weight: make(map[string]float64, 0),
	}
}

func (u *UnionFind) Add(p string) {
	if _, exists := u.parent[p]; exists {
		return
	}
	u.parent[p] = p
	u.weight[p] = 1
}

func (u *UnionFind) Find(p string) string {
	// 用 recursion 的作法可以很輕鬆的在壓縮過程中處理每個節點 weight 的更新 (從最末端開始處理)
	// e.g. a -> b -> c -> d, find(a)
	// recursion 就可以一路指到 d 之後將走過的路徑依序退回來
	if p != u.parent[p] {
		tmp := u.parent[p]
		u.parent[p] = u.Find(u.parent[p])
		u.weight[p] *= u.weight[tmp]
	}
	return u.parent[p]
}

func (u *UnionFind) Union(p, q string, value float64) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
	}

	// 詳細推導請參考 README.md 提供的外部連結 (LeetCodeCN)
	// 懶人包就只是變數求導, let a*b = c*d, 即 d = a*b/c
	u.parent[rootP] = u.parent[rootQ]
	u.weight[rootP] = u.weight[q] * value / u.weight[p]
}

func (u *UnionFind) Query(p, q string) float64 {
	if _, exists := u.parent[p]; !exists {
		return -1
	}
	if _, exists := u.parent[q]; !exists {
		return -1
	}
	// 檢查是不是屬於同一個子集
	if u.Find(p) != u.Find(q) {
		return -1
	}

	// p, q, root 的關係會像這樣 p -> root <- q (i.e. 有共同的 root)
	// 也就是說從 p -> q 的路徑會變成 p -> root -> q
	// 因此 root -> q 的權重就會是 1/u.weight[q]
	return u.weight[p] / u.weight[q]
}
