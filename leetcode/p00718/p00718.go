package p00718

// Time Complexity: O(mn), where m is the number of nums1, n is the number of nums2
// Space Complexity: O(mn)
func findLength(nums1 []int, nums2 []int) int {
	ans := 0
	memo := make([][]int, len(nums1)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(nums2)+1)
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				memo[i+1][j+1] = memo[i][j] + 1
				if memo[i+1][j+1] > ans {
					ans = memo[i+1][j+1]
				}
			}
		}
	}

	return ans
}
