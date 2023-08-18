package p00863

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	result := []int{}

	memo := map[*TreeNode]int{}
	find(root, target, memo)
	dfs(root, target, k, memo[root], memo, &result)

	return result
}

// 找出從 root 到 target 之間的所有 nodes, 並且將 node 與對應的距離記錄到 hash table 內
func find(root, target *TreeNode, memo map[*TreeNode]int) int {
	if root == nil {
		return -1
	}
	if root == target {
		memo[root] = 0
		return 0
	}
	if distance := find(root.Left, target, memo); distance >= 0 {
		memo[root] = distance + 1
		return distance + 1
	}
	if distance := find(root.Right, target, memo); distance >= 0 {
		memo[root] = distance + 1
		return distance + 1
	}
	return -1
}

// 從 root 出發逐一檢查每一個 node 與 target 之間的距離是否剛好等於 k
func dfs(root, target *TreeNode, k, current int, memo map[*TreeNode]int, result *[]int) {
	if root == nil {
		return
	}
	// 關鍵點在這, 每一次檢查 node 跟 target 的距離之前都要先檢查 node 是否曾被 hash table 紀錄過
	// 如果存在的話, 要把 node 的距離更新至正確的
	// 正常來說每往 childs 探尋一次, 就會離 target 越來越遠, 但如果是剛好在 root <-> target 之間的這段距離
	// 每往下探尋一層就表示越接近 target, 因此如果不在比較之前檢查 hash table 的話距離會出現錯誤
	if distance, exists := memo[root]; exists {
		current = distance
	}
	// 找到相等的 node 後不能馬上返回的原因在於, 我們還需要繼續往下找以 target 為主的 subtree 是否有符合條件的 nodes
	if k == current {
		*result = append(*result, root.Val)
	}
	dfs(root.Left, target, k, current+1, memo, result)
	dfs(root.Right, target, k, current+1, memo, result)
}
