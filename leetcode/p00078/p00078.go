package p00078

// Time Complexity: O(2^n), where n is the number of nums
// Space Complexity: O(n)
func subsets(nums []int) [][]int {
	res := [][]int{}
	dfs(&res, []int{}, nums, 0)
	return res
}

func dfs(res *[][]int, path, nums []int, idx int) {
	if idx == len(nums) {
		*res = append(*res, path)
		return
	}
	dfs(res, append([]int{}, path...), nums, idx+1)
	dfs(res, append([]int{}, append(path, nums[idx])...), nums, idx+1)
}
