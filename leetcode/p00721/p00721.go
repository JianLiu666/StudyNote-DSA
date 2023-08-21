package p00721

import "sort"

// Time Complexity: O(n)
// Space Complexity: O(n)
func accountsMerge(accounts [][]string) [][]string {
	size := len(accounts)
	u := CreateUnionFind(size)

	// create union find
	for _, list := range accounts {
		name := list[0]
		rootEmail := list[1]
		u.Insert(name, rootEmail)

		for i := 2; i < len(list); i++ {
			u.Insert(name, list[i])
			// 預設先將第一個 email 當作是 root 把相關的 email 關聯在一起
			u.Union(rootEmail, list[i])
		}
	}

	return u.Result()
}

type UnionFind struct {
	name   map[string]string // mapping email -> name
	parent map[string]string // mapping email -> root email
}

func CreateUnionFind(size int) UnionFind {
	return UnionFind{
		name:   make(map[string]string, size),
		parent: make(map[string]string, size),
	}
}

func (u *UnionFind) Insert(name, email string) {
	if _, exists := u.name[email]; exists {
		return
	}
	u.name[email] = name
	u.parent[email] = email
}

func (u *UnionFind) Find(email string) string {
	for email != u.parent[email] {
		// 為了加速以後查找速度
		u.parent[email] = u.parent[u.parent[email]]
		email = u.parent[email]
	}
	return email
}

func (u *UnionFind) Union(p, q string) {
	rootP, rootQ := u.Find(p), u.Find(q)
	// 以 string 大小做為誰要當作是 root email 的標準
	if rootP < rootQ {
		u.parent[rootQ] = rootP
	} else {
		u.parent[rootP] = rootQ
	}
}

func (u *UnionFind) Result() [][]string {
	set := map[string][]string{}
	for email := range u.parent {
		// 不能直接用現有的 u.parent[email] 當作是 root email
		// 這裡必需特別說明的是, 因為不能保證現在的 group 都已經確保 root email 是最佳化的路徑
		// 所以我們還是得再把現在這個 email 重新找一次確保真正的 rootEmail 是誰
		rootEmail := u.Find(email)
		set[rootEmail] = append(set[rootEmail], email)
	}

	result := [][]string{}
	for rootEmail, emails := range set {
		sort.Strings(emails)
		group := []string{u.name[rootEmail]}
		group = append(group, emails...)
		result = append(result, group)
	}

	return result
}
