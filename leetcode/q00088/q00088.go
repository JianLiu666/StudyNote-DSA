package q00088

// Time Complexity: O(m+n)
// Space Complexity: O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	current := len(nums1) - 1

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[current] = nums1[m-1]
			m--
		} else {
			nums1[current] = nums2[n-1]
			n--
		}
		current--
	}

	for n > 0 {
		nums1[current] = nums2[n-1]
		n--
		current--
	}
}
