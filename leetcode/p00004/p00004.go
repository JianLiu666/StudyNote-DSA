package p00004

// Time Complexity: O(log(m+n)), where m is the number of nums1, n is the number of nums2
// Space Complexity: O(1)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	left := (len1 + len2 + 1) / 2  // 從 1 開始數
	right := (len1 + len2 + 2) / 2 // 從 1 開始數

	return float64(findKth(nums1, 0, len1-1, nums2, 0, len2-1, left)+findKth(nums1, 0, len1-1, nums2, 0, len2-1, right)) * 0.5
}

func findKth(nums1 []int, start1, end1 int, nums2 []int, start2, end2 int, k int) int {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1

	if len1 == 0 {
		return nums2[start2+k-1]
	}
	if len2 == 0 {
		return nums1[start1+k-1]
	}
	if k == 1 {
		return Min(nums1[start1], nums2[start2])
	}

	min1 := start1 + Min(len1, k/2) - 1
	min2 := start2 + Min(len2, k/2) - 1

	if nums1[min1] < nums2[min2] {
		return findKth(nums1, min1+1, end1, nums2, start2, end2, k-(min1-start1+1))
	}
	return findKth(nums1, start1, end1, nums2, min2+1, end2, k-(min2-start2+1))
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
