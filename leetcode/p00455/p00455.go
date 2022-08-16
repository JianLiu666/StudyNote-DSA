package p00455

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func findContentChildren(g []int, s []int) int {
	g = mergesort(g)
	s = mergesort(s)

	gcur, scur := 0, 0
	count := 0
	for gcur < len(g) && scur < len(s) {
		if g[gcur] <= s[scur] {
			count++
			gcur++
		}
		scur++
	}

	return count
}

func mergesort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	larr := mergesort(nums[:mid])
	rarr := mergesort(nums[mid:])

	res := make([]int, len(nums))
	cur := 0
	lcur := 0
	rcur := 0
	for lcur < len(larr) && rcur < len(rarr) {
		if larr[lcur] <= rarr[rcur] {
			res[cur] = larr[lcur]
			lcur++
		} else {
			res[cur] = rarr[rcur]
			rcur++
		}
		cur++
	}
	for lcur < len(larr) {
		res[cur] = larr[lcur]
		lcur++
		cur++
	}
	for rcur < len(rarr) {
		res[cur] = rarr[rcur]
		rcur++
		cur++
	}

	return res
}
