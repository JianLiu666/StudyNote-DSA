package p00839

// Time Complexity: O(n^2), where n is the number of strs
// Space Complexity: O(n)
func numSimilarGroups(strs []string) int {
	u := CreateUnionFind()

	for _, str := range strs {
		u.Add(str)
	}

	// 合併群組
	size := len(strs)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if isSimilar(strs[i], strs[j]) {
				u.Union(strs[i], strs[j])
			}
		}
	}

	// 統計最後區分出來的群組數有多少個
	res := map[string]bool{}
	for _, str := range strs {
		res[u.Find(str)] = true
	}
	return len(res)
}

func isSimilar(str1, str2 string) bool {
	count := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			count++
			if count > 2 {
				return false
			}
		}
	}
	return true
}

type UnionFind struct {
	parents map[string]string
	ranks   map[string]int
}

func CreateUnionFind() UnionFind {
	return UnionFind{
		parents: make(map[string]string, 0),
		ranks:   make(map[string]int, 0),
	}
}

func (u *UnionFind) Add(str string) {
	if _, exists := u.parents[str]; exists {
		return
	}
	u.parents[str] = str
	u.ranks[str] = 1
}

func (u *UnionFind) Find(str string) string {
	for str != u.parents[str] {
		u.parents[str] = u.parents[u.parents[str]]
		str = u.parents[str]
	}
	return str
}

func (u *UnionFind) Union(str1, str2 string) {
	root1 := u.Find(str1)
	root2 := u.Find(str2)
	if root1 == root2 {
		return
	}

	if u.ranks[root2] > u.ranks[root1] {
		u.parents[root1] = root2
		u.ranks[root2] += u.ranks[root1]
	} else {
		u.parents[root2] = root1
		u.ranks[root1] += u.ranks[root2]
	}
}
