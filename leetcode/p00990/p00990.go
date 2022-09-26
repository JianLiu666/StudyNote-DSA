package p00990

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func equationsPossible(equations []string) bool {
	u := CreateUnionFind()
	for _, equation := range equations {
		u.Add(equation[0] - 'a')
		u.Add(equation[3] - 'a')
		if equation[1] == '=' {
			u.Union(equation[0]-'a', equation[3]-'a')
		}
	}

	for _, equation := range equations {
		ans := false
		if equation[1] == '=' {
			ans = true
		}
		if ans != u.IsConnected(equation[0]-'a', equation[3]-'a') {
			return false
		}
	}
	return true
}

type UnionFind struct {
	parents [26]byte
	ranks   [26]int
}

func CreateUnionFind() UnionFind {
	items := [26]byte{}
	for i := 0; i < 26; i++ {
		items[i] = 27
	}

	return UnionFind{
		parents: items,
		ranks:   [26]int{},
	}
}

func (u *UnionFind) Exists(idx byte) bool {
	return u.parents[idx] < 27
}

func (u *UnionFind) Add(idx byte) {
	if u.Exists(idx) {
		return
	}
	u.parents[idx] = idx
}

func (u *UnionFind) Find(idx byte) byte {
	if idx != u.parents[idx] {
		u.parents[idx] = u.Find(u.parents[idx])
	}
	return u.parents[idx]
}

func (u *UnionFind) IsConnected(idx1, idx2 byte) bool {
	root1 := u.Find(idx1)
	root2 := u.Find(idx2)
	return root1 == root2
}

func (u *UnionFind) Union(idx1, idx2 byte) {
	root1 := u.Find(idx1)
	root2 := u.Find(idx2)
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
