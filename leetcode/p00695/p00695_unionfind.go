package p00695

import "fmt"

// Time Complexity: O(mn), where m is the height of grid, n is the width of grid
// Space Complexity: O(mn)
func maxAreaOfIsland_UnionFind(grid [][]int) int {
	u := CreateUnionFind()

	for row, columns := range grid {
		for col, val := range columns {
			if val == 1 {
				key := genKey(row, col)
				u.Add(key)

				if row-1 >= 0 && grid[row-1][col] == 1 {
					u.Union(key, genKey(row-1, col))
				}
				if col-1 >= 0 && grid[row][col-1] == 1 {
					u.Union(key, genKey(row, col-1))
				}
			}
		}
	}

	return u.Maximum()
}

func genKey(row, col int) string {
	return fmt.Sprintf("%v,%v", row, col)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type UnionFind struct {
	parent  map[string]string
	rank    map[string]int
	maximum int
}

func CreateUnionFind() UnionFind {
	return UnionFind{
		parent:  make(map[string]string, 0),
		rank:    make(map[string]int, 0),
		maximum: 0,
	}
}

func (u *UnionFind) Maximum() int {
	return u.maximum
}

func (u *UnionFind) Add(key string) {
	u.parent[key] = key
	u.rank[key] = 1
	u.maximum = max(u.maximum, 1)
}

func (u *UnionFind) Find(key string) string {
	for key != u.parent[key] {
		u.parent[key] = u.parent[u.parent[key]]
		key = u.parent[key]
	}

	return key
}

func (u *UnionFind) Union(p, q string) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
	}

	if u.rank[rootQ] > u.rank[rootP] {
		u.parent[rootP] = rootQ
		u.rank[rootQ] += u.rank[rootP]
		u.maximum = max(u.maximum, u.rank[rootQ])
	} else {
		u.parent[rootQ] = rootP
		u.rank[rootP] += u.rank[rootQ]
		u.maximum = max(u.maximum, u.rank[rootP])
	}
}
