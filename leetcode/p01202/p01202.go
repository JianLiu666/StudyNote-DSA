package p01202

import (
	"sort"
	"strings"
)

// Time Complexity: O(nlogn), where n is the length of s
// Space Complexity: O(n), where n is the length of s
func smallestStringWithSwaps(s string, pairs [][]int) string {
	// union find 關聯子集合
	u := CreateUnionFind(len(s))
	// tc: O(n)
	for _, pair := range pairs {
		// tc: O(1)
		u.Union(pair[0], pair[1])
	}

	// 把同一個子集合的字母整理在一起
	memo := map[int][]rune{}
	// tc: O(n)
	for i, ch := range s {
		group := u.Find(i)
		if memo[group] == nil {
			memo[group] = make([]rune, 0)
		}
		memo[group] = append(memo[group], ch)
	}

	// 重新排序同一個子集合的字母順序
	// tc: O(n)
	for _, arr := range memo {
		// tc: O(nlogn)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	}

	// 按照題目給定的 s pattern 重新輸出排序後的結果
	count := map[int]int{}
	var str strings.Builder
	// tc: O(n)
	for i := range s {
		group := u.Find(i)
		str.WriteRune(memo[group][count[group]])
		count[group]++
	}

	return str.String()
}

type UnionFind struct {
	parent []int
	rank   []int
}

func CreateUnionFind(size int) UnionFind {
	instance := UnionFind{
		parent: make([]int, size),
		rank:   make([]int, size),
	}

	for i := 0; i < size; i++ {
		instance.parent[i] = i
		instance.rank[i] = 1
	}

	return instance
}

func (u *UnionFind) Find(p int) int {
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}
	return p
}

func (u *UnionFind) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
	}

	if u.rank[rootQ] > u.rank[rootP] {
		u.parent[rootP] = rootQ
	} else {
		u.parent[rootQ] = rootP
		if u.rank[rootP] == u.rank[rootQ] {
			u.rank[rootP]++
		}
	}
}
