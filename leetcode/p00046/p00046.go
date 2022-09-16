package p00046

// Time Complexity: O(n!)
// Space Complexity: O(n)
func permute(nums []int) [][]int {
	res := [][]int{}
	dfs(&res, []int{}, nums)
	return res
}

func dfs(res *[][]int, path, nums []int) {
	if len(nums) == 1 {
		*res = append(*res, append(append([]int{}, path...), nums...))
		return
	}

	for i := 0; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(res, path, append(append([]int{}, nums[:i]...), nums[i+1:]...))
		path = path[:len(path)-1]
	}
}
